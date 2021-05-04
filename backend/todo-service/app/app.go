// Package app provides application level functionality.
package app

import (
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-service/app/config"
	"todo-service/app/handlers/todos"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

// App struct
type App struct {
	config   *config.Config
	logger   *zerolog.Logger
	db       *sql.DB
	natsConn *nats.Conn
}

// New func
func New(c *config.Config, l *zerolog.Logger) *App {
	return &App{config: c, logger: l}
}

// Run method
func (a *App) Run() {

	a.logInfo("application is starting up...")

	if err := a.initDB(); err != nil {
		a.logFatal(err)
	}

	if err := a.initNATS(); err != nil {
		a.logFatal(err)
	}

	tc := todos.NewController(
		todos.NewSQLRepository(a.db, a.logger),
		a.natsConn,
		a.logger)

	if err := tc.HandleTodoCreated(); err != nil {
		a.logFatal(err)
	}

	a.logInfo("application is running")

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	<-s

	a.logInfo("application is shutting down...")

	a.close()

	a.logInfo("application is shut down")
}

func (a *App) initDB() error {

	db, err := sql.Open("pgx", a.config.DBConnectionString)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return err
	}

	a.db = db

	return nil
}

func (a *App) initNATS() error {

	nc, err := nats.Connect(
		a.config.NATSUrl,
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(10),
		nats.ReconnectWait(time.Second),
	)

	if err != nil {
		return err
	}

	a.natsConn = nc

	return nil
}

func (a *App) close() {
	_ = a.natsConn.Drain()
	_ = a.db.Close()
}

func (a *App) logInfo(msg string) {
	a.logger.Info().Msg(msg)
}

func (a *App) logFatal(err error) {
	a.logger.Fatal().Err(err).Send()
}
