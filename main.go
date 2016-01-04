package main

import (
	"github.com/astaxie/beego"
	_ "workSummary/routers"
)

func main() {
	beego.SessionOn = true
	beego.Run()
}
