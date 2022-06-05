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
	Config   *Config
	Logger   *zerolog.Logger
	DB       *sql.DB
	NATSConn *nats.Conn
	Sub      *NATSMessageSubscriber
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

	s.initMessageSubscriber()

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

func (s *Service) initMessageSubscriber() {
	parser := &NATSMessageParser{NewSchemaValidator(schemaFS)}
	repo := &PostgresRepository{s.DB}
	pub := NewCentrifugoMessagePublisher(s.Config)

	options := &NATSMessageSubscriberOptions{
		Subject:   "todo.*",
		Durable:   "todo",
		BatchSize: 1,
		Handlers: map[string]NATSMessageHandler{
			"todo.create":   &TodoCreateHandler{parser, repo, pub, s.Logger},
			"todo.complete": &TodoCompleteHandler{parser, repo, pub, s.Logger},
		},
	}

	s.Sub = &NATSMessageSubscriber{options, s.NATSConn, s.Logger}
}

func (s *Service) serve(ctx context.Context) error {
	errChan := make(chan error)

	go func() {
		<-ctx.Done()

		s.Logger.Info().Msgf("application is shutting down...")

		_ = s.NATSConn.Drain()
		_ = s.DB.Close()

		errChan <- nil
	}()

	if err := s.Sub.Subscribe(ctx); err != nil {
		return err
	}

	s.Logger.Info().Msgf("application is running")

	return <-errChan
}
