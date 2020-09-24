package routers

import (
	"beego1/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//新建一个路径/register
    beego.Router("/register", &controllers.MaintoController{})
}
