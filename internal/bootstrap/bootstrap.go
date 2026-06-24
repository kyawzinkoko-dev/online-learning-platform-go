package bootstrap

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kyawzinkoko-dev/online-learning-platform/configs"
	"github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/auth/handler"
	authRouter "github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/auth/route"
	"github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/auth/service"
	"github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/user/repository"
	"github.com/kyawzinkoko-dev/online-learning-platform/pkg/database"
)

type Application struct {
	Engine *gin.Engine
	Config *configs.Config
}

func Build() (*Application, error) {
	cfg, err := configs.LoadConfig()

	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := database.NewPostgresDatabae(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	userRepo := repository.NewPostgresRepository(db)
	authService := service.NewAuthService(userRepo, cfg)
	authHandler := handler.NewAuthHandler(authService)

	apiV1 := engine.Group("/api/v1")

	authRouter.AuthRoute(apiV1, authHandler)

	return &Application{
		Engine: engine,
		Config: cfg,
	}, nil

}

func (app *Application) Run() error {
	log.Printf("Starting server on port %s", app.Config.AppPort)
	return app.Engine.Run(":" + app.Config.AppPort)
}
