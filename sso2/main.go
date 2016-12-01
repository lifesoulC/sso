package main

import (
	"sso2/g"
	_ "sso2/routers"

	"github.com/astaxie/beego"
)

func main() {
	g.InitEnv()
	beego.Run()
}
