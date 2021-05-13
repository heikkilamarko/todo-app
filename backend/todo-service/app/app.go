// Package app provides application level functionality.
package app

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-service/app/config"
	"todo-service/app/todos"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

// App struct
type App struct {
	config *config.Config
	logger *zerolog.Logger
	db     *sql.DB
	nc     *nats.Conn
	js     nats.JetStreamContext
}

// New func
func New(c *config.Config, l *zerolog.Logger) *App {
	return &App{config: c, logger: l}
}

// Run method
func (a *App) Run() {

	a.logInfo("application is starting up...")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := a.initDB(ctx); err != nil {
		a.logFatal(err)
	}

	if err := a.initNATS(); err != nil {
		a.logFatal(err)
	}

	if err := a.registerRoutes(ctx); err != nil {
		a.logFatal(err)
	}

	if err := a.serve(ctx); err != nil {
		a.logFatal(err)
	}

	a.logInfo("application is shut down")
}

func (a *App) initDB(ctx context.Context) error {

	db, err := sql.Open("pgx", a.config.DBConnectionString)
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

	a.db = db

	return nil
}

func (a *App) initNATS() error {

	nc, err := nats.Connect(
		a.config.NATSUrl,
		nats.Token(a.config.NATSToken),
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(10),
		nats.ReconnectWait(time.Second),
	)

	if err != nil {
		return err
	}

	js, err := nc.JetStream()

	if err != nil {
		return err
	}

	a.nc = nc
	a.js = js

	return nil
}

func (a *App) registerRoutes(ctx context.Context) error {

	c := todos.NewController(a.config, a.logger, a.db, a.nc, a.js)

	if err := c.Start(ctx); err != nil {
		return err
	}

	return nil
}

func (a *App) serve(ctx context.Context) error {

	var (
		err = make(chan error)
	)

	go func() {
		<-ctx.Done()

		a.logInfo("application is shutting down...")

		_ = a.nc.Drain()
		_ = a.db.Close()

		err <- nil
	}()

	a.logInfo("application is running")

	return <-err
}

func (a *App) logInfo(msg string, v ...interface{}) {
	a.logger.Info().Msgf(msg, v...)
}

func (a *App) logFatal(err error) {
	a.logger.Fatal().Err(err).Send()
}
