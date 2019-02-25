package routers

import (
	"qwe/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/index", &controllers.MainController{},"post:Index")
}
