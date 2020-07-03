package routers

import (
	"OA/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/loadmeun", &controllers.MainController{}, "post:LoadMeun")

	beego.Router("/action", &controllers.ActionController{}, "get:Index")
	beego.Router("/action/add", &controllers.ActionController{}, "get,post:Add")
	beego.Router("/action/getlist", &controllers.ActionController{}, "post:GetList")

	beego.Router("/role", &controllers.RoleController{}, "get:Index")
	beego.Router("/role/add", &controllers.RoleController{}, "get,post:Add")
	beego.Router("/role/getlist", &controllers.RoleController{}, "post:GetList")
}
