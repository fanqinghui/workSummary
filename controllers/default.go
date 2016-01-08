package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type ToRegisterController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["SystemName"] = beego.AppConfig.String("systemName")
	c.Data["ApplicationName"] = beego.AppConfig.String("applicationName")
	c.Data["CompanyName"] = beego.AppConfig.String("companyName")
	c.TplNames = "login.html"
}
