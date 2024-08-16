package routes

import (
	"event-app/controllers"
	"event-app/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "event-app/docs" // Import generated docs
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.POST("/events", controllers.CreateEvent)
		auth.GET("/events", controllers.GetEvents)
		auth.GET("/events/today", controllers.GetTodayEvents)
	}

	return r
}
