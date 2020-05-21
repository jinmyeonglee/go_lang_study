package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func show_indx_page(c * gin.Context) {
	c.HTML(
		http.StatusOK,

	)
}
