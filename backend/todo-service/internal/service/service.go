package service

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-service/internal/adapters"
	"todo-service/internal/application"
	"todo-service/internal/application/command"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

type config struct {
	App                string `json:"app"`
	DBConnectionString string `json:"db_connection_string"`
	NATSUrl            string `json:"nats_url"`
	NATSToken          string `json:"nats_token"`
	LogLevel           string `json:"log_level"`
}

type Service struct {
	config                *config
	logger                *zerolog.Logger
	db                    *sql.DB
	nc                    *nats.Conn
	js                    nats.JetStreamContext
	app                   *application.Application
	todoMessageSubscriber *adapters.TodoNATSMessageSubscriber
}

func (s *Service) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s.loadConfig()
	s.initLogger()

	s.logInfo("application is starting up...")

	if err := s.initDB(ctx); err != nil {
		s.logFatal(err)
	}

	if err := s.initNATS(); err != nil {
		s.logFatal(err)
	}

	s.initApplication()
	s.initMessageSubscriber()

	if err := s.serve(ctx); err != nil {
		s.logFatal(err)
	}

	s.logInfo("application is shut down")
}

func (s *Service) loadConfig() {
	s.config = &config{
		App:                env("APP_NAME", ""),
		DBConnectionString: env("APP_DB_CONNECTION_STRING", ""),
		NATSUrl:            env("APP_NATS_URL", ""),
		NATSToken:          env("APP_NATS_TOKEN", ""),
		LogLevel:           env("APP_LOG_LEVEL", "warn"),
	}
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
		s.config.NATSUrl,
		nats.Token(s.config.NATSToken),
		nats.NoReconnect(),
		nats.DisconnectErrHandler(
			func(_ *nats.Conn, err error) {
				s.logFatal(err)
			}),
		nats.ErrorHandler(
			func(_ *nats.Conn, _ *nats.Subscription, err error) {
				s.logFatal(err)
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

func (s *Service) initApplication() {
	todoRepository := adapters.NewTodoPostgresRepository(s.db)
	todoMessagePublisher := adapters.NewTodoNATSMessagePublisher(s.nc)

	s.app = &application.Application{
		Commands: application.Commands{
			CreateTodo:   command.NewCreateTodoHandler(todoRepository, todoMessagePublisher),
			CompleteTodo: command.NewCompleteTodoHandler(todoRepository, todoMessagePublisher),
		},
	}
}

func (s *Service) initMessageSubscriber() {
	s.todoMessageSubscriber = adapters.NewTodoNATSMessageSubscriber(s.app, s.js, s.logger)
}

func (s *Service) serve(ctx context.Context) error {
	errChan := make(chan error)

	go func() {
		<-ctx.Done()

		s.logInfo("application is shutting down...")

		_ = s.nc.Drain()
		_ = s.db.Close()

		errChan <- nil
	}()

	if err := s.todoMessageSubscriber.Subscribe(ctx); err != nil {
		return err
	}

	s.logInfo("application is running")

	return <-errChan
}
