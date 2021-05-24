package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowHomepage response the homepage
func ShowHomepage(c *gin.Context) {
	c.HTML(http.StatusOK, "frame.tmpl", gin.H{})
}

// ShowToolBox response the toolbox column
func ShowToolBox(c *gin.Context) {
	c.HTML(http.StatusOK, "frame.tmpl", gin.H{})
}

// ShowProcessOnline response the toolbox column
func ShowProcessOnline(c *gin.Context) {
	c.HTML(http.StatusOK, "frame.tmpl", gin.H{})
}

// ShowDataShare response the toolbox column
func ShowDataShare(c *gin.Context) {
	c.HTML(http.StatusOK, "frame.tmpl", gin.H{})
}

// ShowContactUs response the toolbox column
func ShowContactUs(c *gin.Context) {
	c.HTML(http.StatusOK, "frame.tmpl", gin.H{})
}

// ShowUser response the toolbox column
func ShowUser(c *gin.Context) {
	c.HTML(http.StatusOK, "frame.tmpl", gin.H{})
}
