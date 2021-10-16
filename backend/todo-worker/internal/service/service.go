package service

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-worker/internal/workflow"

	"github.com/rs/zerolog"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	zerologadapter "logur.dev/adapter/zerolog"
	"logur.dev/logur"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

type config struct {
	App                string
	DBConnectionString string
	TemporalHostPort   string
	LogLevel           string
}

type Service struct {
	config *config
	logger *zerolog.Logger
	db     *sql.DB
	tc     client.Client
	tw     worker.Worker
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

	if err := s.initTemporal(); err != nil {
		s.logFatal(err)
	}

	if err := s.serve(ctx); err != nil {
		s.logFatal(err)
	}

	s.logInfo("application is shut down")
}

func (s *Service) loadConfig() {
	s.config = &config{
		App:                env("APP_NAME", ""),
		DBConnectionString: env("APP_DB_CONNECTION_STRING", ""),
		TemporalHostPort:   env("APP_TEMPORAL_HOSTPORT", ""),
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

func (s *Service) initTemporal() error {
	tc, err := client.NewClient(client.Options{
		HostPort: s.config.TemporalHostPort,
		Logger:   logur.LoggerToKV(zerologadapter.New(*s.logger)),
	})

	if err != nil {
		return err
	}

	a := workflow.NewActivities(s.db)

	tw := worker.New(tc, workflow.TaskQueueWorker, worker.Options{})
	tw.RegisterWorkflow(workflow.RemoveTodosWorkflow)
	tw.RegisterActivity(a.RemoveTodos)

	s.tc = tc
	s.tw = tw

	return nil
}

func (s *Service) serve(ctx context.Context) error {
	errChan := make(chan error)

	go func() {
		<-ctx.Done()

		s.logInfo("application is shutting down...")

		_ = s.db.Close()
		s.tc.Close()

		errChan <- nil
	}()

	s.logInfo("application is running")

	if err := s.tw.Run(worker.InterruptCh()); err != nil {
		return err
	}

	return <-errChan
}
