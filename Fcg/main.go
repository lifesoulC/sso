package main

import (
	"Fcg/g"
	_ "Fcg/routers"

	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	g.InitEnv()
	err := g.Loginit()
	if err != nil {
		fmt.Println("loginit err")
	}
	beego.Run()
}
