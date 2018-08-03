package routers

import (
	"ctf/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/UserInfo", &controllers.UserInfoController{})
}
