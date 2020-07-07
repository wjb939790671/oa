package routers

import (
	"OA/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/loadmeun", &controllers.MainController{}, "post:LoadMeun")

	beego.Router("/action", &controllers.ActionController{}, "get:Index")
	beego.Router("/action/index", &controllers.ActionController{}, "get:Index")
	beego.Router("/action/add", &controllers.ActionController{}, "get,post:Add")
	beego.Router("/action/getlist", &controllers.ActionController{}, "post:GetList")
<<<<<<< HEAD
	beego.Router("/action/updata", &controllers.RoleController{}, "get,post:Updata")
	beego.Router("/action/delete", &controllers.RoleController{}, "get,post:Delete")
=======
	beego.Router("/action/updata", &controllers.ActionController{}, "get,post:Updata")
	beego.Router("/action/delete", &controllers.ActionController{}, "get,post:Delete")
	beego.Router("/action/query", &controllers.ActionController{}, "post:QueryAll")
>>>>>>> 8a1f30b2a36b34695cf2098c09d2fbb2e1fcb52f

	beego.Router("/role", &controllers.RoleController{}, "get:Index")
	beego.Router("/role/index", &controllers.RoleController{}, "get:Index")
	beego.Router("/role/add", &controllers.RoleController{}, "get,post:Add")
	beego.Router("/role/getlist", &controllers.RoleController{}, "post:GetList")
	beego.Router("/role/updata", &controllers.RoleController{}, "get,post:Updata")
	beego.Router("/role/delete", &controllers.RoleController{}, "get,post:Delete")

	beego.Router("/employee", &controllers.EmployeeController{}, "get:Index")
	beego.Router("/employee/index", &controllers.EmployeeController{}, "get:Index")
	beego.Router("/employee/add", &controllers.EmployeeController{}, "get,post:Add")
	beego.Router("/employee/getlist", &controllers.EmployeeController{}, "post:GetList")
	beego.Router("/employee/updata", &controllers.EmployeeController{}, "get,post:Updata")
	beego.Router("/employee/delete", &controllers.EmployeeController{}, "get,post:Delete")
}
