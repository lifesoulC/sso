package routers

import (
	"sso2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Default")
	beego.Router("/login", &controllers.LoginController{}, "get:Login;post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	beego.Router("/add", &controllers.RuleController{}, "get:AddRule;post:DoAddRule")
	beego.Router("/show", &controllers.RuleController{}, "get:ShowDB")
	beego.Router("/delete", &controllers.RuleController{}, "get:DeleteRule")
	beego.Router("/edit", &controllers.RuleController{}, "get:EditRule;post:DoEditRule")
	beego.Router("/goto", &controllers.RuleController{}, "get:GoToPage")
	//beego.Router("/nodelink", &controllers.RuleController{}, "get:SetNodeLink;post:DoNodeLink")
	beego.Router("/addrole", &controllers.RuleController{}, "get:AddRole;post:DoAddRole")

}
