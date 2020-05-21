package main
import (
	"github.com/gin-gonic/gin"
	"router/init_routes.go"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("pulic/templates/*")

	initialize_routes()

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
