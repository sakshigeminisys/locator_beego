package routers

import (
	"locator/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello-world/:id([0-9]+)", &controllers.MainController{}, "get:HelloSitepoint")
    beego.Router("/nodes", &controllers.NodeController{}, "get:GetNodes")
    beego.Router("/getservice", &controllers.ServiceController{}, "get:GetService")
    beego.Router("/service/:id", &controllers.ServiceController{}, "get:Service")
    beego.Router("/register", &controllers.ServiceController{}, "put,post:Register")
    beego.Router("/deregister/:id([0-9]+)", &controllers.ServiceController{}, "put,post:Deregister")

}
