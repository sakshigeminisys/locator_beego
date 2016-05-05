package routers

import (
	"locator/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/nodes", &controllers.ServiceController{}, "get:GetNodes")
    beego.Router("/getservice", &controllers.ServiceController{}, "get:GetService")
    beego.Router("/service/:id", &controllers.ServiceController{}, "get:Service")
    beego.Router("/register", &controllers.ServiceController{}, "put,post:Register")
    beego.Router("/deregister/:id([0-9]+)", &controllers.ServiceController{}, "put,post:Deregister")

}
