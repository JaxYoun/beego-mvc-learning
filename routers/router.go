package routers

import (
	"beego-mvc-learning/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test0", &controllers.MainController{})
	beego.Router("/postTest", &controllers.MainController{})
}
