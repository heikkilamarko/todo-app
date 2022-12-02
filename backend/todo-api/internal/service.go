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
	Config   *Config
	Logger   *zerolog.Logger
	DB       *sql.DB
	NATSConn *nats.Conn
	Server   *http.Server
}

func (s *Service) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := s.loadConfig(); err != nil {
		s.Logger.Fatal().Err(err).Send()
	}

	s.initLogger()

	s.Logger.Info().Msgf("application is starting up...")

	if err := s.initDB(ctx); err != nil {
		s.Logger.Fatal().Err(err).Send()
	}

	if err := s.initNATS(); err != nil {
		s.Logger.Fatal().Err(err).Send()
	}

	s.initHTTPServer(ctx)

	if err := s.serve(ctx); err != nil {
		s.Logger.Fatal().Err(err).Send()
	}

	s.Logger.Info().Msgf("application is shut down")
}

func (s *Service) loadConfig() error {
	c := &Config{}
	if err := c.Load(); err != nil {
		return err
	}

	s.Config = c

	return nil
}

func (s *Service) initLogger() {
	level, err := zerolog.ParseLevel(s.Config.LogLevel)
	if err != nil {
		level = zerolog.WarnLevel
	}

	zerolog.SetGlobalLevel(level)

	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Str("app", s.Config.App).
		Logger()

	s.Logger = &logger
}

func (s *Service) initDB(ctx context.Context) error {
	db, err := sql.Open("pgx", s.Config.DBConnectionString)
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

	s.DB = db

	return nil
}

func (s *Service) initNATS() error {
	conn, err := nats.Connect(
		s.Config.NATSURL,
		nats.Token(s.Config.NATSToken),
		nats.NoReconnect(),
		nats.DisconnectErrHandler(
			func(_ *nats.Conn, err error) {
				s.Logger.Fatal().Err(err).Send()
			}),
		nats.ErrorHandler(
			func(_ *nats.Conn, _ *nats.Subscription, err error) {
				s.Logger.Fatal().Err(err).Send()
			}),
	)

	if err != nil {
		return err
	}

	s.NATSConn = conn

	return nil
}

func (s *Service) initHTTPServer(ctx context.Context) {
	router := mux.NewRouter()

	jwtConfig := &middleware.JWTConfig{
		Issuer:   s.Config.AuthIssuer,
		Iss:      s.Config.AuthClaimIss,
		Aud:      []string{s.Config.AuthClaimAud},
		TokenKey: ContextKeyAccessToken,
	}

	router.Use(
		middleware.Logger(s.Logger),
		middleware.RequestLogger(),
		middleware.ErrorRecovery(),
		middleware.JWT(ctx, jwtConfig),
	)

	repo := &PostgresRepository{s.DB}
	pub := &NATSMessagePublisher{s.NATSConn}

	router.Handle("/todos/userinfo", &GetUserinfoHandler{repo, s.Logger}).Methods(http.MethodGet)
	router.Handle("/todos/token", &GetCentrifugoTokenHandler{s.Config, s.Logger}).Methods(http.MethodGet)
	router.Handle("/todos", &GetTodosHandler{repo, s.Logger}).Methods(http.MethodGet)
	router.Handle("/todos", &CreateTodoHandler{pub, s.Logger}).Methods(http.MethodPost)
	router.Handle("/todos/{id:[0-9]+}/complete", &CompleteTodoHandler{pub, s.Logger}).Methods(http.MethodPost)

	router.NotFoundHandler = goutils.NotFoundHandler()

	s.Server = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         s.Config.Address,
		Handler:      router,
	}
}

func (s *Service) serve(ctx context.Context) error {
	errChan := make(chan error)

	go func() {
		<-ctx.Done()

		s.Logger.Info().Msgf("application is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = s.Server.Shutdown(ctx)
		_ = s.NATSConn.Drain()
		_ = s.DB.Close()

		errChan <- nil
	}()

	s.Logger.Info().Msgf("application is running at %s", s.Server.Addr)

	if err := s.Server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return <-errChan
}
