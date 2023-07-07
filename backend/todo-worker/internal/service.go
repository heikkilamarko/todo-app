package internal

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Service struct {
	Config *Config
	Logger *slog.Logger
	DB     *sql.DB
	Client client.Client
	Worker worker.Worker
}

func (s *Service) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := s.loadConfig(); err != nil {
		s.Logger.Error(err.Error())
		os.Exit(1)
	}

	s.initLogger()

	s.Logger.Info("application is starting up...")

	if err := s.initDB(ctx); err != nil {
		s.Logger.Error(err.Error())
		os.Exit(1)
	}

	if err := s.initTemporal(); err != nil {
		s.Logger.Error(err.Error())
		os.Exit(1)
	}

	if err := s.serve(ctx); err != nil {
		s.Logger.Error(err.Error())
		os.Exit(1)
	}

	s.Logger.Info("application is shut down")
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
	level := slog.LevelInfo

	level.UnmarshalText([]byte(s.Config.LogLevel))

	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler := slog.NewJSONHandler(os.Stderr, opts).
		WithAttrs([]slog.Attr{
			slog.String("app", s.Config.App),
		})

	logger := slog.New(handler)

	slog.SetDefault(logger)

	s.Logger = logger
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
	tc, err := client.Dial(client.Options{
		HostPort: s.Config.TemporalHostPort,
		Logger:   s.Logger,
	})

	if err != nil {
		return err
	}

	a := Activities{s.DB}

	tw := worker.New(tc, TaskQueue, worker.Options{})
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

		s.Logger.Info("application is shutting down...")

		_ = s.DB.Close()
		s.Client.Close()

		errChan <- nil
	}()

	s.Logger.Info("application is running")

	if err := s.Worker.Run(worker.InterruptCh()); err != nil {
		return err
	}

	return <-errChan
}
