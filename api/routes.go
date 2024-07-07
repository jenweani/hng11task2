package api

import (
	"os"
	"hng11task2/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func BuildRoutesHandler() *gin.Engine {
	r := gin.New()

	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET("/health", handlers.HealthHandler)

	// Auth routes
	authRoute := r.Group("/auth")
	authRoute.POST("/register", handlers.RegisterUserHandler)
	authRoute.POST("/login")

	apiRoutes := r.Group("/api")

	apiRoutes.POST("/organisations/:orgId/users", handlers.AddUserToOrganisation)
	apiRoutes.Use(AuthMiddleware())

	apiRoutes.GET("/users/:id", handlers.GetUserByID)
	apiRoutes.GET("/organisations", handlers.GetAllOrgsForSignedInUser)
	apiRoutes.GET("/organisations/:orgId")
	apiRoutes.POST("/organisations", handlers.CreateOrganisation)
	



	// User routes

	// Organisation routes

	return r
}