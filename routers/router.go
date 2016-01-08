package routers

import (
	"workSummary/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/logout", &controllers.LogoutController{}, "get:Logout")
	beego.Router("/register", &controllers.RegisterController{}, "post:Register")
	beego.Router("/toRegister", &controllers.ToRegisterController{}, "get:ToRegister")
	beego.Router("/todayWork", &controllers.TodayWorkController{}, "get:TodayWork")
	beego.Router("/saveTodayWork", &controllers.SaveTodayWorkController{}, "post:SaveTodayWork")
}
