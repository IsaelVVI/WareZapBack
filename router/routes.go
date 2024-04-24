package router

import (
	"github.com/IsaelVVI/warezapback.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// initialize Handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/tests", handler.ListOpeningHandler)

		v1.GET("/teste", handler.ShowOpeningHandler)

		v1.POST("/teste", handler.CreateOpeningHandler)

		v1.PUT("/teste", handler.UpdateOpeningHandler)

		v1.DELETE("/teste", handler.DeleteOpeningHandler)

	}
}
