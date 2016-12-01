package controllers

//"fmt"

type AdminController struct {
	BaseController
}

func (this *AdminController) CheckLogin() {
	if !this.IsAdmin {
		this.Redirect("http://sso.lbase.inc", 302)
	}
}
