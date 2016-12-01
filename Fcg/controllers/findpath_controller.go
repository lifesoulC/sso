package controllers

import (
	"Fcg/g"
	//"encoding/json"
	"fmt"
	//"os"
	//	"strconv"
	"strings"
)

//type LbRoute struct {
//	SrcAddr string
//	SrcIp   string

//	dstAddr string
//	dstip   string
//}

//type Nodesmpinfo struct {
//	NodeId   string `json:"nodeId"` //"1-1"
//	Location string `json:"location"`
//	Ip       string `json:"ip"`
//	Isp      string `json:"isp"` //"北京"
//}
//type RoutlinkMap struct {
//	From   Nodesmpinfo `json:"from"`   //源地址
//	To     Nodesmpinfo `json:"to"`     //目的地址
//	Level  int         `json:"level"`  //浏量等级
//	Delays int         `json:"delay"`  //延迟  "20ms"
//	Lost   string      `json:"lost"`   //丢包率
//	Status int         `json:"status"` //默认为0表示自动生成的  1 为手动修改
//}

func (this *NodeController) FindPath() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	type Data struct {
		Net   []g.RoutelinkMapNet `json:"net"`
		Nodes []g.RoutlinkMap     `json:"nodes"`
	}
	type Response struct {
		Status int  `json:"status"`
		Data   Data `json:"data"`
	}

	response := Response{}
	fromid := this.GetString("from")
	toid := this.GetString("to")
	level := this.GetString("level")
	fmt.Println("fromid", fromid)
	fmt.Println("toid", toid)
	fmt.Println("level", level)

	g.Routeing(fromid, toid, level)
	fmt.Println("findpath 52", g.LbRoutetmp)
	if len(g.LbRoutetmp) == 0 {
		response.Status = 0
		response.Data.Net = nil
		response.Data.Nodes = nil
		this.Data["json"] = response
		this.ServeJSON()
		return
	}

	for _, routetmp := range g.LbRoutetmp {
		s := strings.Split(routetmp.SrcAddr, "-") //解析发来的字符
		srcNetId := s[0]
		srcNodeId := s[1]

		dsts := strings.Split(routetmp.DstAddr, "-")
		dstNetId := dsts[0]
		//dstNodeId := dsts[1]

		if srcNetId == dstNetId {
			tmp := g.RouteMap[srcNetId][srcNodeId][routetmp.DstAddr]
			response.Data.Nodes = append(response.Data.Nodes, tmp)
		} else {
			tmp := g.RouteMapNet[srcNetId][srcNodeId][routetmp.DstAddr]
			response.Data.Net = append(response.Data.Net, tmp)
		}

	}
	g.LbRoutetmp = nil
	response.Status = 1
	this.Data["json"] = response
	this.ServeJSON()

}
