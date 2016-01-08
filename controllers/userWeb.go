package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"workSummary/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

//全局session
//var globalSessions *session.Manager
/*
func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()
}*/

type LoginController struct {
	beego.Controller
}
type LogoutController struct {
	beego.Controller
}
type RegisterController struct {
	beego.Controller
}
type TodayWorkController struct {
	beego.Controller
}
type SaveTodayWorkController struct {
	beego.Controller
}

/**
登录
**/
func (this *LoginController) Login() {

	fmt.Println("Login")
	loginName := this.GetString("loginName")
	password := this.GetString("password")
	//校验
	valid := validation.Validation{}
	valid.Required(loginName, "loginName")
	valid.Required(password, "password")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	fmt.Println(loginName, password)

	user, flag := models.Login(loginName, password)
	if flag {
		this.SetSession("User", user) //存放session
		this.Data["User"] = user
		this.TplNames = "index.html"
	} else {
		this.Data["SystemName"] = beego.AppConfig.String("systemName")
		this.Data["ApplicationName"] = beego.AppConfig.String("applicationName")
		this.Data["CompanyName"] = beego.AppConfig.String("companyName")
		this.Data["LoginName"] = loginName
		this.Data["Password"] = password
		this.Data["Message"] = "用户名密码错误"
		this.TplNames = "login.html"
	}

}

/**
退出
**/
func (this *LogoutController) Logout() {
	user := this.GetSession("User") //存放session
	if u, ok := user.(models.User); ok {
		fmt.Println(u.LoginName, u.LoginName, u.Id)
	}

	this.DelSession("User") //删除session
	this.TplNames = "login.html"
}

/**
跳转到注册页面
**/
func (c *ToRegisterController) ToRegister() {
	c.TplNames = "register.html"
}

/**
注册
**/
func (this *RegisterController) Register() {
	fmt.Println("register")
	loginName := this.GetString("loginName")
	password := this.GetString("password")
	userName := this.GetString("userName")
	email := this.GetString("email")
	//校验
	valid := validation.Validation{}
	valid.Required(userName, "username")
	valid.Required(password, "password")
	valid.Required(loginName, "loginName")
	valid.Required(email, "email")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	user := models.User{
		LoginName: loginName,
		UserName:  userName,
		Password:  password,
		Email:     email,
	}
	fmt.Println(loginName, password, userName, email)

	user, flag := models.InsertUser(user)
	if flag {
		this.Data["Message"] = "注册成功，请登录"
	} else {
		this.Data["Message"] = "注册失败，请联系管理员"
	}
	this.Data["SystemName"] = beego.AppConfig.String("systemName")
	this.Data["ApplicationName"] = beego.AppConfig.String("applicationName")
	this.Data["CompanyName"] = beego.AppConfig.String("companyName")
	this.TplNames = "register.html"
}

/**
今日工作
**/
func (this *TodayWorkController) TodayWork() {
	//查询当前用户的日报
	nowTime := time.Now()
	fmt.Println(nowTime.Year(), nowTime.Month(), nowTime.Day())
	dailyWork, ok := models.GetTodayWork(nowTime.Year(), int(nowTime.Month()), nowTime.Day())
	if ok {
		fmt.Println(dailyWork)
		this.Data["ContentId"] = dailyWork.Id
		this.Data["Content"] = dailyWork.Content
	} else {
		fmt.Println("无今日日报")
	}

	this.TplNames = "todayWork.html"
}

/**
今日工作
**/
func (this *SaveTodayWorkController) SaveTodayWork() {
	fmt.Println("SaveTodayWorkController")
	content := this.GetString("content")
	contentId, _ := this.GetInt("contentId")
	fmt.Println(contentId, content)

	user := this.GetSession("User")
	if u, ok := user.(models.User); ok {

		nowTime := time.Now()
		todayWork := models.DailyWork{
			Id:      contentId,
			Content: content,
			Year:    nowTime.Year(),
			Month:   int(nowTime.Month()),
			Day:     nowTime.Day(),
			UId:     u.Id,
		}

		returnJson := models.SaveTodayWork(todayWork)

		if key, err := json.Marshal(returnJson); err == nil {
			fmt.Println(string(key))
			this.Data["json"] = string(key)
		}
	} else {
		fmt.Println("user is null")
	}
	this.ServeJson()
}

/**
以下是登录校验
**/
func (this *TodayWorkController) Prepare() {
	fmt.Println("todayWorkController Prepare ing ==========================")
	user := this.GetSession("User")
	if _, ok := user.(models.User); !ok {
		this.Redirect("/", 302)
	}

}

func (this *LogoutController) Prepare() {
	fmt.Println("LogoutController Prepare ing ==========================")
	user := this.GetSession("User")
	if _, ok := user.(models.User); !ok {
		this.Redirect("/", 302)
	}
}

func (this *SaveTodayWorkController) Prepare() {
	fmt.Println("SaveTodayWorkController Prepare ing ==========================")
	user := this.GetSession("User")
	if _, ok := user.(models.User); !ok {
		this.Redirect("/", 302)
	}
}
