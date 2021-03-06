package server

import (
	"embed"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/api/openapi"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo/http/rest"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo/repository/postgresql"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo/service"
)

// Config defines the required configurations by the server.
type Config struct {
	Address  string
	DB       *pgxpool.Pool
	Logger   *zap.Logger
	Static   embed.FS
	Validate *validator.Validate
}

// New returns a configured server that contains all internal API codebase.
func New(cfg Config) (*http.Server, error) {
	// Construct new gin with default options.
	r := gin.Default()

	// Set gin mode that depends on ENV.
	gin.SetMode(os.Getenv("GIN_MODE"))

	// Register ginzap middleware.
	r.Use(ginzap.Ginzap(cfg.Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(cfg.Logger, true))

	// Implement cors middleware to gin.
	r.Use(cors.Default())

	// Register repositories used for dependency injection to the services.
	todoRepo := postgresql.NewTodo(cfg.DB)

	// Register services used for dependency injection to the handlers.
	todoSvc := service.NewTodo(cfg.Logger, todoRepo)

	// Serve swagger specification files.
	openapi.RegisterSpecifications(r)

	// Serve embedded static files that live on a rest-server/static folder.
	r.GET("/static/*filepath", func(c *gin.Context) {
		c.FileFromFS(c.Request.URL.Path, http.FS(cfg.Static))
	})

	// Register all REST HTTP handlers.
	rest.NewTodoHandler(cfg.Validate, todoSvc).Register(r)

	return &http.Server{
		Addr:    cfg.Address,
		Handler: r,
	}, nil
}
