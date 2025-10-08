package internal

import (
	"database/sql"
	"log"
	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/middleware"
	"user-service/internal/repository"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
)

type App struct {
	config *config.Config
	router *gin.Engine
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {
	cfg := config.Load()
	a.config = cfg

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, cfg.JWTSecret)
	userHandler := handler.NewUserHandler(userService)

	a.setupRouter(userHandler)

	log.Printf("start user service on port %s", cfg.Port)
	return a.router.Run(":" + cfg.Port)
}

func (a *App) setupRouter(h *handler.UserHandler) {
	a.router = gin.Default()

	a.router.Use(middleware.CORS())
	a.router.Use(gin.Logger())

	api := a.router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/register", h.Register)
			users.POST("/login", h.Login)
			users.GET("/:id", h.GetUser)
		}
	}
}
