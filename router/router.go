package router

// import do gin - framework http GO lang
import "github.com/gin-gonic/gin"

func Initialize() {
	// geralmente as variaveis são efemeras - curtas

	// Init routers with gin framework
	router := gin.Default()

	// Init routes
	initializeRoutes(router)

	// Run the Application/Server in a port your like
	router.Run(":3000") // listen and serve on 0.0.0.0:3000
}
