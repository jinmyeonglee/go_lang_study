package routers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func show_indx_page(c * gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi~!",
	})
}
