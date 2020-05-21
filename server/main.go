package main
import (
	"github.com/gin-gonic/gin"
	"../routers"
)
// Fix import package problem

var router *gin.Engine

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("pulic/templates/*")

	initialize_routes()
	println("asdf")
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
