package routers

import (
	"goweb/dataServer/controller"

	"github.com/gin-gonic/gin"
)

// Setup function initialize router engine
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("templates/frame.tmpl")
	r.Static("/static", "./static")
	LoadFrame(r)
	return r
}

// loadFrame is the function load the html frame
func LoadFrame(r *gin.Engine) {
	r.GET("/", controller.HomepageController)
	r.GET("/toolbox", controller.ToolboxController)
	r.GET("/processonline", controller.ProcessonlineController)
	r.GET("/datashare", controller.DatashareController)
	r.GET("/contactus", controller.ContactusController)
	r.GET("/login", controller.LoginController)
	r.GET("/logout", controller.LogoutController)
	r.GET("/changepw", controller.ChangepwController)
	r.GET("/admin", controller.AdminController)
}
