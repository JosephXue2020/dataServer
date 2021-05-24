package routers

import (
	"goweb/dataServer/controller"

	"github.com/gin-gonic/gin"
)

// Init function initialize router engine
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("templates/frame.tmpl")
	r.Static("/static", "./static")
	LoadFrame(r)
	return r
}

// loadFrame is the function load the html frame
func LoadFrame(r *gin.Engine) {
	r.GET("", controller.ShowHomepage)
	r.GET("/toolbox", controller.ShowToolBox)
	r.GET("/processonline", controller.ShowProcessOnline)
	r.GET("/datashare", controller.ShowDataShare)
	r.GET("/contactus", controller.ShowContactUs)
	r.GET("/user", controller.ShowUser)
}
