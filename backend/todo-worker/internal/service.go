package internal

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	zerologadapter "logur.dev/adapter/zerolog"
	"logur.dev/logur"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Service struct {
	Config *Config
	Logger *zerolog.Logger
	DB     *sql.DB
	Client client.Client
	Worker worker.Worker
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

	if err := s.initTemporal(); err != nil {
		s.Logger.Fatal().Err(err).Send()
	}

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

func (s *Service) initTemporal() error {
	tc, err := client.NewClient(client.Options{
		HostPort: s.Config.TemporalHostPort,
		Logger:   logur.LoggerToKV(zerologadapter.New(*s.Logger)),
	})

	if err != nil {
		return err
	}

	a := Activities{s.DB}

	tw := worker.New(tc, TaskQueueWorker, worker.Options{})
	tw.RegisterWorkflow(RemoveTodosWorkflow)
	tw.RegisterActivity(a.RemoveTodos)

	s.Client = tc
	s.Worker = tw

	return nil
}

func (s *Service) serve(ctx context.Context) error {
	errChan := make(chan error)

	go func() {
		<-ctx.Done()

		s.Logger.Info().Msgf("application is shutting down...")

		_ = s.DB.Close()
		s.Client.Close()

		errChan <- nil
	}()

	s.Logger.Info().Msgf("application is running")

	if err := s.Worker.Run(worker.InterruptCh()); err != nil {
		return err
	}

	return <-errChan
}
