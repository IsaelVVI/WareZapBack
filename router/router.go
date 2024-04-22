package router

// import do gin - framework http GO lang
import "github.com/gin-gonic/gin"

func Initialize() {
	// geralmente as variaveis são efemeras - curtas
	// Inicia o Router com as configs padroes do GIN
	router := gin.Default()

	// Definição de rotas!
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Rodando a APLICAÇÃO
	router.Run(":3000") // listen and serve on 0.0.0.0:3000
}
