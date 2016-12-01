package controllers

import (
	"Fcg/g"
	"encoding/json"
	"fmt"
	//"os"
	//	"strconv"
	//"strings"
)

func (this *NodeController) Lbroute() {

	sendroute := g.SendRoute{}

	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &sendroute)
	if err != nil {
		fmt.Println(" json error", err)
	}

	//	s := strings.Split(sendroute.Addr, "-") //解析发来的字符
	//	NetId := s[0]
	//	NodeId := s[1]

	if sendroute.Type == "node" {
		g.LbRouteNode[sendroute.Addr] = string(sendroute.Content)

	}

	if sendroute.Type == "net" {
		g.LbRouteNet[sendroute.Addr] = string(sendroute.Content)
	}

	fmt.Println("lb_controller 35 sendroute:", sendroute.Type, string(sendroute.Content))
	this.Ctx.WriteString("")

}
