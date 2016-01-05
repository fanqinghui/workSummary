package routers

import (
	"workSummary/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/toRegister", &controllers.ToRegisterController{})
	beego.Router("/todayWork", &controllers.TodayWorkController{})
	beego.Router("/saveTodayWork", &controllers.SaveTodayWorkController{})

}
