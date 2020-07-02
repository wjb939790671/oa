package routers

import (
	"OA/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/loadmeun", &controllers.MainController{}, "post:LoadMeun")
}
