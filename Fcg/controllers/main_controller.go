package controllers

type MainController struct {
	AdminController
}

func (this *MainController) Get() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	//	type oo struct {
	//		Ip     string `json:"ip"`
	//		Status int    `json:"status"`
	//		Isp    string `json:"isp"`
	//	}
	//	op := oo{}
	//	op.Ip = "sdfdsf"
	//	op.Status = 789
	//	op.Isp = "poi"
	//	this.Data["Website"] = "beego.me"

	//	this.Data["Email"] = "astaxie@gmail.com"

	//	this.Data["json"] = op

	//	this.ServeJSON()
	this.TplName = "index.html"
}
