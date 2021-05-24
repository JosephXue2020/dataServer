package routers

import (
	"goweb/dataServer/controller"

	"github.com/gin-gonic/gin"
)

// Init function initialize router engine
func Init() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("templates/frame.html")
	LoadFrame(r)
	return r
}

// loadFrame is the function load the html frame
func LoadFrame(r *gin.Engine) {
	r.GET("", controller.ShowHomepage)
}
