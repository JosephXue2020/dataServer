package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowHomepage(c *gin.Context) {
	c.HTML(http.StatusOK, "templates/frame.html", gin.H{})
}
