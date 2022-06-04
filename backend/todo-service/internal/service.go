package internal

import (
	"context"
	"database/sql"
	"embed"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

//go:embed schemas/*.json
var schemaFS embed.FS

type Service struct {
	config *Config
	logger *zerolog.Logger
	db     *sql.DB
	nc     *nats.Conn
	sub    *NATSMessageSubscriber
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

	s.initMessageSubscriber()

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

	s.nc = nc

	return nil
}

func (s *Service) initMessageSubscriber() {
	parser := &NATSMessageParser{NewSchemaValidator(schemaFS)}
	repo := &PostgresRepository{s.db}
	pub := NewCentrifugoMessagePublisher(s.config)

	options := &NATSMessageSubscriberOptions{
		Subject:   "todo.*",
		Durable:   "todo",
		BatchSize: 1,
		Handlers: map[string]NATSMessageHandler{
			"todo.create":   &TodoCreateHandler{parser, repo, pub, s.logger},
			"todo.complete": &TodoCompleteHandler{parser, repo, pub, s.logger},
		},
	}

	s.sub = &NATSMessageSubscriber{options, s.nc, s.logger}
}

func (s *Service) serve(ctx context.Context) error {
	errChan := make(chan error)

	go func() {
		<-ctx.Done()

		s.logger.Info().Msgf("application is shutting down...")

		_ = s.nc.Drain()
		_ = s.db.Close()

		errChan <- nil
	}()

	if err := s.sub.Subscribe(ctx); err != nil {
		return err
	}

	s.logger.Info().Msgf("application is running")

	return <-errChan
}
