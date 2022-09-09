package routes

import (
	"authentication_service/cmd/handlers"
	"authentication_service/cmd/repository"
	"authentication_service/cmd/usecase"
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(conn *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Repository
	authenticationRepository := repository.NewRepositoryUser(conn)
	// Service
	authenticationUsecase := usecase.NewUseCaseUser(authenticationRepository)
	// Handler
	authenticationHandlers := handlers.NewHandler(authenticationUsecase)

	// Prefix
	api := router.Group("/api/v1")
	// Endpoint register
	api.POST("/register", authenticationHandlers.Register)

	return router
}
