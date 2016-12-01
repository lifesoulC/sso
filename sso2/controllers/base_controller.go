package controllers

import (
	"sso2/models"

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
	name := this.Ctx.GetCookie("name")
	password := this.Ctx.GetCookie("password")
	if name == "" || password == "" {
		this.IsAdmin = false
		return
	}
	person, err := models.ExtractOnePersonById(name)
	if err != nil {
		this.Ctx.WriteString("用户名不正确")
		return
	} else {
		if password != person.Password {
			this.Ctx.WriteString("密码不正确")
			return
		}
	}

	this.IsAdmin = true
	this.Data["IsAdmin"] = this.IsAdmin

}
