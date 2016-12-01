package controllers

import (
	"Fcg/g"
	"encoding/json"
	"fmt"
	//"os"
	//	"strconv"
	"strings"
)

type NodeController struct {
	AdminController
}

func (this *NodeController) Nodeinit() { //初始化接受过来的node信息 初始化NodeMap map[string]map[string]Nodemap
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	Nodemapinfo := g.Nodelinkmap{}
	sendinfo := g.Sendinfo{}

	confginfo := g.Confg{}
	//	ips := g.IPs{}
	//	publicIp := g.IpMsg{}
	//	publicIp.Ip = "123.123.123.123"
	//	publicIp.Isp = "联通"
	//	ip1 := "10.0.0.1"
	//	ip2 := "10.0.0.3"
	//	ips.PrivateIp = append(ips.PrivateIp, ip1, ip2)
	//	ips.PublicIp = append(ips.PublicIp, publicIp)

	//	confginfo.Addr = "1-1"
	//	confginfo.Name = "gz09"
	//	confginfo.Location = "北京"
	//	confginfo.HostType = "vm"
	//	confginfo.Ip = "15.15.15.15"

	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &sendinfo)
	if err != nil {
		fmt.Println(" json error", err)
	}
	fmt.Println("node_contrller 42 sendinco:", sendinfo)
	confginfo = sendinfo.Confginfo

	s := strings.Split(confginfo.Addr, "-") //解析发来的字符
	NetId := s[0]
	NodeId := s[1]

	for _, v := range sendinfo.Ips.PublicIp {
		tmpipinfo := g.IPinfo{}
		tmpipinfo.Ip = v.Ip
		tmpipinfo.Isp = v.Isp
		tmpipinfo.Status = 1
		Nodemapinfo.Ippublic = append(Nodemapinfo.Ippublic, tmpipinfo)

	}

	for _, v := range sendinfo.Ips.PrivateIp {
		tmpipinfo := g.IPinfo{}
		tmpipinfo.Ip = v
		tmpipinfo.Isp = ""
		tmpipinfo.Status = 1
		Nodemapinfo.Ipprivate = append(Nodemapinfo.Ipprivate, tmpipinfo)
	}

	Nodemapinfo.Flow = ""
	Nodemapinfo.Id = confginfo.Addr
	Nodemapinfo.Name = confginfo.Name
	Nodemapinfo.Location = confginfo.Location
	Nodemapinfo.HostType = confginfo.HostType
	Nodemapinfo.Net = NetId
	Nodemapinfo.Node = NodeId
	Nodemapinfo.Status = 1
	Nodemapinfo.Ip = confginfo.Ip
	//fmt.Println(Nodemapinfo)

	fmt.Println("node_controller 76 Nodemapinfo:", Nodemapinfo)
	g.AddNodemapinfo(Nodemapinfo)
	for node, _ := range g.NodeMap[Nodemapinfo.Net] {
		g.CompRout(Nodemapinfo.Net, node) //计算出cfg
		fmt.Println("node_controller 81 RoutnodeMap:", g.RouteMap)
		g.InitCfg(Nodemapinfo.Net) //下发配置
	}

	//	g.Node

	//g.RoutNetMap[routemsg.Addr] = string(routemsg.Content)
	//fmt.Println("82 :", g.RouteMap)
	this.Ctx.WriteString("")
}
