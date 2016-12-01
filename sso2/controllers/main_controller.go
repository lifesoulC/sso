package controllers

type MainController struct {
	AdminController
}

func (this *MainController) Default() {
	this.Layout = "layout/admin.html"
	this.TplName = "index.html"
}
