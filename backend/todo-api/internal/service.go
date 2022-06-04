package internal

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
	"github.com/heikkilamarko/goutils/middleware"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Service struct {
	config *Config
	logger *zerolog.Logger
	db     *sql.DB
	nc     *nats.Conn
	js     nats.JetStreamContext
	server *http.Server
}

func (s *Service) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := s.loadConfig(); err != nil {
		s.logger.Fatal().Err(err).Send()
	}

	s.initLogger()

	s.logger.Info().Msgf("application is starting up...")

	if err := s.initDB(ctx); err != nil {
		s.logger.Fatal().Err(err).Send()
	}

	if err := s.initNATS(); err != nil {
		s.logger.Fatal().Err(err).Send()
	}

	s.initHTTPServer(ctx)

	if err := s.serve(ctx); err != nil {
		s.logger.Fatal().Err(err).Send()
	}

	s.logger.Info().Msgf("application is shut down")
}

func (s *Service) loadConfig() error {
	c := &Config{}
	if err := c.Load(); err != nil {
		return err
	}

	s.config = c

	return nil
}

func (s *Service) initLogger() {
	level, err := zerolog.ParseLevel(s.config.LogLevel)
	if err != nil {
		level = zerolog.WarnLevel
	}

	zerolog.SetGlobalLevel(level)

	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Str("app", s.config.App).
		Logger()

	s.logger = &logger
}

func (s *Service) initDB(ctx context.Context) error {
	db, err := sql.Open("pgx", s.config.DBConnectionString)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.PingContext(ctx); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Service) initNATS() error {
	nc, err := nats.Connect(
		s.config.NATSURL,
		nats.Token(s.config.NATSToken),
		nats.NoReconnect(),
		nats.DisconnectErrHandler(
			func(_ *nats.Conn, err error) {
				s.logger.Fatal().Err(err).Send()
			}),
		nats.ErrorHandler(
			func(_ *nats.Conn, _ *nats.Subscription, err error) {
				s.logger.Fatal().Err(err).Send()
			}),
	)

	if err != nil {
		return err
	}

	js, err := nc.JetStream()

	if err != nil {
		return err
	}

	s.nc = nc
	s.js = js

	return nil
}

func (s *Service) initHTTPServer(ctx context.Context) {
	router := mux.NewRouter()

	jwtConfig := &middleware.JWTConfig{
		Issuer:   s.config.AuthIssuer,
		Iss:      s.config.AuthClaimIss,
		Aud:      []string{s.config.AuthClaimAud},
		TokenKey: ContextKeyAccessToken,
	}

	router.Use(
		middleware.Logger(s.logger),
		middleware.RequestLogger(),
		middleware.ErrorRecovery(),
		middleware.JWT(ctx, jwtConfig),
	)

	repo := NewPostgresTodoRepository(s.db)
	pub := NewNATSTodoMessagePublisher(s.js)

	router.Handle("/todos/token", &GetCentrifugoTokenHandler{s.config, s.logger}).Methods(http.MethodGet)
	router.Handle("/todos", &GetTodosHandler{repo, s.logger}).Methods(http.MethodGet)
	router.Handle("/todos", &CreateTodoHandler{pub, s.logger}).Methods(http.MethodPost)
	router.Handle("/todos/{id:[0-9]+}/complete", &CompleteTodoHandler{pub, s.logger}).Methods(http.MethodPost)

	router.NotFoundHandler = goutils.NotFoundHandler()

	s.server = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         s.config.Address,
		Handler:      router,
	}
}

func (s *Service) serve(ctx context.Context) error {
	errChan := make(chan error)

	go func() {
		<-ctx.Done()

		s.logger.Info().Msgf("application is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = s.server.Shutdown(ctx)
		_ = s.nc.Drain()
		_ = s.db.Close()

		errChan <- nil
	}()

	s.logger.Info().Msgf("application is running at %s", s.server.Addr)

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return <-errChan
}
