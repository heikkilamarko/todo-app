// Package service ...
package service

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-api/adapters"
	"todo-api/app"
	"todo-api/app/command"
	"todo-api/app/query"
	"todo-api/ports"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
	"github.com/heikkilamarko/goutils/middleware"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

type config struct {
	App                string `json:"app"`
	Address            string `json:"address"`
	DBConnectionString string `json:"db_connection_string"`
	NATSUrl            string `json:"nats_url"`
	NATSToken          string `json:"nats_token"`
	LogLevel           string `json:"log_level"`
}

type Service struct {
	config     *config
	logger     *zerolog.Logger
	db         *sql.DB
	nc         *nats.Conn
	js         nats.JetStreamContext
	app        *app.App
	httpServer *ports.HTTPServer
	router     *mux.Router
	server     *http.Server
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

	s.initApp()
	s.initRouter()
	s.initHTTPServer()
	s.initServer()

	if err := s.serve(ctx); err != nil {
		s.logFatal(err)
	}

	s.logInfo("application is shut down")
}

func (s *Service) loadConfig() {
	s.config = &config{
		App:                env("APP_NAME", ""),
		Address:            env("APP_ADDRESS", ":8080"),
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

func (s *Service) initApp() {
	todoRepo := adapters.NewTodoRepository(s.db, s.js)

	s.app = &app.App{
		Commands: app.Commands{
			CreateTodo:   command.NewCreateTodoHandler(todoRepo),
			CompleteTodo: command.NewCompleteTodoHandler(todoRepo),
		},
		Queries: app.Queries{
			GetTodos: query.NewGetTodosHandler(todoRepo),
		},
	}
}

func (s *Service) initRouter() {
	router := mux.NewRouter()

	router.Use(
		middleware.Logger(s.logger),
		middleware.RequestLogger(),
		middleware.ErrorRecovery(),
	)

	router.NotFoundHandler = goutils.NotFoundHandler()

	s.router = router
}

func (s *Service) initHTTPServer() {
	s.httpServer = ports.NewHTTPServer(s.app, s.router, s.logger)
}

func (s *Service) initServer() {
	s.server = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         s.config.Address,
		Handler:      s.router,
	}
}

func (s *Service) serve(ctx context.Context) error {
	errChan := make(chan error)

	go func() {
		<-ctx.Done()

		s.logInfo("application is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = s.server.Shutdown(ctx)
		_ = s.nc.Drain()
		_ = s.db.Close()

		errChan <- nil
	}()

	s.logInfo("application is running at %s", s.server.Addr)

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return <-errChan
}
