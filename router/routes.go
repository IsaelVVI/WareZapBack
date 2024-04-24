package router

import (
	"github.com/IsaelVVI/warezapback.git/docs"
	"github.com/IsaelVVI/warezapback.git/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// initialize Handler
	handler.InitializeHandler()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/tests", handler.ListOpeningHandler)

		v1.GET("/teste", handler.ShowOpeningHandler)

		v1.POST("/teste", handler.CreateOpeningHandler)

		v1.PUT("/teste", handler.UpdateOpeningHandler)

		v1.DELETE("/teste", handler.DeleteOpeningHandler)

	}
	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
