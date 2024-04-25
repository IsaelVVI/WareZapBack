package router

import (
	"github.com/IsaelVVI/warezapback.git/controllers"
	user_controller "github.com/IsaelVVI/warezapback.git/controllers/users"
	"github.com/IsaelVVI/warezapback.git/docs"
	"github.com/IsaelVVI/warezapback.git/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// initialize Controller
	controllers.InitializeHandler()

	basePath := "/api/v1"
	usersPath := basePath + "/user"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath, middlewares.Auth())
	{
		// Test routes
		v1.GET("/tests", controllers.ListOpening)

		v1.GET("/teste", controllers.ShowOpening)

		v1.POST("/teste", controllers.CreateOpening)

		v1.PUT("/teste", controllers.UpdateOpening)

		v1.DELETE("/teste", controllers.DeleteOpening)

	}

	user := router.Group(usersPath)
	{
		user.POST("/create", user_controller.CreatedUser)
	}

	auth := router.Group(basePath + "/auth")
	{
		auth.POST("/login", user_controller.Login)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
