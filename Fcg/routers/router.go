package routers

import (
	"Fcg/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/nodeinit", &controllers.NodeController{}, "post:Nodeinit")
	beego.Router("/delay", &controllers.NodeController{}, "post:Delay")

	beego.Router("/poll", &controllers.NodeController{}, "post:Poll")
	beego.Router("/lbroute", &controllers.NodeController{}, "post:Lbroute")

	beego.Router("/nodeInfo", &controllers.NodeController{}, "get:NodeInfo")
	beego.Router("/nodeOperator/status/", &controllers.NodeController{}, "post:NodeOperatorStatus")
	beego.Router("/nodeOperator/location/", &controllers.NodeController{}, "post:NodeOperatorLocation")
	beego.Router("/nodeOperator/ip/", &controllers.NodeController{}, "post:NodeOperatorIp")
	beego.Router("/findPath", &controllers.NodeController{}, "post:FindPath")
	beego.Router("/showLbRoute", &controllers.NodeController{}, "post:ShowLbRoute")

	beego.Router("/path", &controllers.NodeController{}, "get:Path")
	beego.Router("/changePath", &controllers.NodeController{}, "post:ChangePath")
}
