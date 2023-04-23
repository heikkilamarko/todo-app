package internal

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nats-io/nats.go"
	"golang.org/x/exp/slog"

	// PostgreSQL driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Service struct {
	Config   *Config
	Logger   *slog.Logger
	DB       *sql.DB
	Repo     Repository
	Pub      MessagePublisher
	AuthZ    AuthZ
	NATSConn *nats.Conn
	Server   *http.Server
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

	if err := s.initAuthZ(ctx); err != nil {
		s.Logger.Error(err.Error())
		os.Exit(1)
	}

	if err := s.initNATS(); err != nil {
		s.Logger.Error(err.Error())
		os.Exit(1)
	}

	s.initHTTPServer(ctx)

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

	opts := slog.HandlerOptions{
		Level: level,
	}

	handler := opts.NewJSONHandler(os.Stderr).
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
	s.Repo = &PostgresRepository{s.DB}

	return nil
}

func (s *Service) initAuthZ(ctx context.Context) error {
	s.Logger.Info("init authz", "authz_backend", s.Config.AuthZBackend)

	switch s.Config.AuthZBackend {
	case "db":
		s.AuthZ = NewDBAuthZ(s.Repo)
	default:
		authZ, err := NewOPAAuthZ(ctx)
		if err != nil {
			return err
		}
		s.AuthZ = authZ
	}

	return nil
}

func (s *Service) initNATS() error {
	conn, err := nats.Connect(
		s.Config.NATSURL,
		nats.Token(s.Config.NATSToken),
		nats.MaxReconnects(-1),
		nats.DisconnectErrHandler(
			func(_ *nats.Conn, err error) {
				s.Logger.Info("nats disconnected", "reason", err)
			}),
		nats.ReconnectHandler(
			func(c *nats.Conn) {
				s.Logger.Info("nats reconnected", "address", c.ConnectedUrl())
			}),
		nats.ErrorHandler(
			func(_ *nats.Conn, _ *nats.Subscription, err error) {
				s.Logger.Error("nats error", "err", err)
				os.Exit(1)
			}),
	)

	if err != nil {
		return err
	}

	s.NATSConn = conn
	s.Pub = &NATSMessagePublisher{s.NATSConn}

	return nil
}

func (s *Service) initHTTPServer(ctx context.Context) {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(s.jwt(ctx))

	router.
		With(s.authZ(ctx, "todo.read")).
		Method(http.MethodGet, "/todos/userinfo", &GetUserinfoHandler{s.Repo, s.Logger})

	router.
		With(s.authZ(ctx, "todo.read")).
		Method(http.MethodGet, "/todos/token", &GetCentrifugoTokenHandler{s.Config, s.Logger})

	router.
		With(s.authZ(ctx, "todo.read")).
		Method(http.MethodGet, "/todos", &GetTodosHandler{s.Repo, s.Logger})

	router.
		With(s.authZ(ctx, "todo.write")).
		Method(http.MethodPost, "/todos", &CreateTodoHandler{s.Pub, s.Logger})

	router.
		With(s.authZ(ctx, "todo.write")).
		Method(http.MethodPost, "/todos/{id:[0-9]+}/complete", &CompleteTodoHandler{s.Pub, s.Logger})

	router.NotFound(NotFound)

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

		s.Logger.Info("application is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = s.Server.Shutdown(ctx)
		_ = s.NATSConn.Drain()
		_ = s.DB.Close()

		errChan <- nil
	}()

	s.Logger.Info("application is running", "port", s.Server.Addr)

	if err := s.Server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return <-errChan
}

func (s *Service) jwt(ctx context.Context) func(http.Handler) http.Handler {
	c := JWTMiddlewareConfig{
		Issuer:     s.Config.AuthIssuer,
		Iss:        s.Config.AuthClaimIss,
		Aud:        []string{s.Config.AuthClaimAud},
		ContextKey: ContextKeyAccessToken,
		Logger:     s.Logger,
	}

	return JWTMiddleware(ctx, &c)
}

func (s *Service) authZ(ctx context.Context, permission string) func(http.Handler) http.Handler {
	c := AuthZMiddlewareConfig{
		AuthZ:      s.AuthZ,
		Permission: permission,
		ContextKey: ContextKeyAuthZResult,
		Logger:     s.Logger,
	}

	return AuthZMiddleware(ctx, &c)
}
