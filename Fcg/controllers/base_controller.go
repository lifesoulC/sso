package controllers

import (
	"github.com/astaxie/beego"
)

type Checker interface {
	CheckLogin()
}

type BaseController struct {
	beego.Controller
	IsAdmin bool
}

func (this *BaseController) Prepare() {
	if app, ok := this.AppController.(Checker); ok {
		this.AssignIsAdmin()
		app.CheckLogin()
	}
}

func (this *BaseController) AssignIsAdmin() {
	test := this.Ctx.GetCookie("test")
	//password := this.Ctx.GetCookie("password")
	//fmt.Println("hello", test)
	if test == "" {
		this.IsAdmin = false
		return
	}
	//fmt.Println(test)
	if test != "John" {
		this.IsAdmin = false
		return
	}

	this.IsAdmin = true
	this.Data["IsAdmin"] = this.IsAdmin

}

//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}
