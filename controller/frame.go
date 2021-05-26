package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ColField is struct type of each column
type Link struct {
	Name string
	Url  string
}

// SetColumn function setup columns of the web
func SetColumn() []Link {
	sli := []Link{
		{"首页", "/"},
		{"工具箱", "/toolbox"},
		{"数据在线处理", "/processonline"},
		{"数据共享", "/datashare"},
		{"联系我们", "/contactus"},
	}
	return sli
}

// enviroment varialbes
var Column []Link

// SetDropdown function setup dropdown columns of the web
func SetDropdown() map[string]interface{} {
	m := make(map[string]interface{})
	m["toggle"] = Link{"用户", "#"}
	m["menu"] = []Link{
		{"登录", "/login"},
		{"登出", "/logout"},
		{"修改密码", "/changepw"},
		{"管理员", "/admin"},
	}
	return m
}

// enviroment varialbes
var Dropdown map[string]interface{}

// SetFootbar function setup footbar of the web
func SetFootbar() []Link {
	sli := make([]Link, 0)
	sli = append(sli, Link{"百科社讯网", "http://itc.ecph.com.cn/"})
	sli = append(sli, Link{"百科三版资源库", "http://192.168.12.129:8080/discovery/site/search"})
	return sli
}

// enviroment variables
var Footbar []Link

// Element variables
var Element map[string]interface{}

// init function
func init() {
	Column = SetColumn()
	Dropdown = SetDropdown()
	Footbar = SetFootbar()
	Element = map[string]interface{}{
		"column":   Column,
		"dropdown": Dropdown,
		"footbar":  Footbar,
	}
}

// HomepageController response the homepage
func HomepageController(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.tmpl", Element)
}

func ToolboxController(c *gin.Context) {
	c.HTML(http.StatusOK, "toolbox.tmpl", Element)
}

func ProcessonlineController(c *gin.Context) {
	c.HTML(http.StatusOK, "processonline.tmpl", Element)
}

func DatashareController(c *gin.Context) {
	c.HTML(http.StatusOK, "datashare.tmpl", Element)
}

func ContactusController(c *gin.Context) {
	c.HTML(http.StatusOK, "contactus.tmpl", Element)
}

func UserController(c *gin.Context) {
	c.HTML(http.StatusOK, "frame.tmpl", Element)
}

func LoginController(c *gin.Context) {
	c.HTML(http.StatusOK, "frame.tmpl", Element)
}

func LogoutController(c *gin.Context) {
	c.HTML(http.StatusOK, "frame.tmpl", Element)
}

func ChangepwController(c *gin.Context) {
	c.HTML(http.StatusOK, "frame.tmpl", Element)
}

func AdminController(c *gin.Context) {
	c.HTML(http.StatusOK, "frame.tmpl", Element)
}
