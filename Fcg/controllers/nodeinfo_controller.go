package controllers

import (
	"Fcg/g"
	//"encoding/json"
	"fmt"
	"time"
	//"os"
	//	"strconv"
	"strings"
)

func (this *NodeController) NodeInfo() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	datas := []g.Data{}
	response := g.Response{}
	//NodeMap = make(map[string]map[string]Nodelinkmap)

	for _, nodemsg := range g.NodeMap {
		for _, nodemap := range nodemsg {
			data := g.Data{}
			data.Flow = nodemap.Flow
			data.Id = nodemap.Id

			tmpipinfo := []g.IPinfo{}
			for _, ipinfo := range nodemap.Ippublic {
				tmpipinfo = append(tmpipinfo, ipinfo)
			}
			for _, ipinfo := range nodemap.Ipprivate {
				tmpipinfo = append(tmpipinfo, ipinfo)
			}
			data.Ipinfo = tmpipinfo

			data.Name = nodemap.Name
			data.Location = nodemap.Location
			data.Net = nodemap.Net
			data.Status = nodemap.Status

			datas = append(datas, data)
		}
	}

	if datas == nil {
		response.Status = 0
	} else {
		response.Status = 1
	}
	response.Datas = datas

	this.Data["json"] = response
	this.ServeJSON()
}

func (this *NodeController) NodeOperatorStatus() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	type Data struct {
		nodeId string `json:"nodeId"`
	}
	type Response struct {
		Status int    `json:"status"`
		Data   string `json:data` //返回错误信息
	}
	response := Response{}
	nodeId := this.GetString("nodeId")
	types := this.GetString("type")
	fmt.Println(nodeId)
	s := strings.Split(nodeId, "-") //解析发来的字符
	NetId := s[0]
	NodeId := s[1]
	if tmpnodemap, ok := g.NodeMap[NetId][NodeId]; ok {
		if types == "stop" {
			tmpnodemap.Status = 0
			g.NodeMap[NetId][NodeId] = tmpnodemap
			g.NilCfg(NetId, NodeId)
		} else {
			tmpnodemap.Status = 1
			g.NodeMap[NetId][NodeId] = tmpnodemap
		}
		time.Sleep(3 * time.Second)
		nodemap := g.NodeMap[NetId]

		for nodeid, _ := range nodemap {
			g.CompRout(NetId, nodeid) //计算出cfg 下发
			g.InitCfg(NetId)
		}

		response.Status = 1
		response.Data = "success"
		this.Data["json"] = response
		this.ServeJSON()

		//	g.InitCfg(NetId)

	} else {
		fmt.Println("Key Not Found")
		response.Status = 0
		response.Data = "not fund " + nodeId
		this.Data["json"] = response
		this.ServeJSON()
	}

}
func (this *NodeController) NodeOperatorLocation() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")

	type Response struct {
		Status int    `json:"status"`
		Data   string `json:data` //返回错误信息
	}
	nodeId := this.GetString("nodeId")
	location := this.GetString("location")

	response := Response{}

	s := strings.Split(nodeId, "-") //解析发来的字符
	NetId := s[0]
	NodeId := s[1]
	if tmpnodemap, ok := g.NodeMap[NetId][NodeId]; ok {
		tmpnodemap.Location = location
		g.NodeMap[NetId][NodeId] = tmpnodemap
		response.Status = 1
		response.Data = "success"
		this.Data["json"] = response
		this.ServeJSON()

	} else {
		fmt.Println("Key Not Found")
		response.Status = 0
		response.Data = "not fund " + nodeId
		this.Data["json"] = response
		this.ServeJSON()
	}

}

func (this *NodeController) NodeOperatorIp() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	type Response struct {
		Status int    `json:"status"`
		Data   string `json:"data"` //返回错误信息
	}
	response := Response{}

	nodeId := this.GetString("nodeId")
	types := this.GetString("type")
	ip := this.GetString("ip")

	s := strings.Split(nodeId, "-") //解析发来的字符
	NetId := s[0]
	NodeId := s[1]
	if tmpnodemap, ok := g.NodeMap[NetId][NodeId]; ok {
		for id, publicip := range tmpnodemap.Ippublic {
			if publicip.Ip == ip {
				if types == "start" {
					publicip.Status = 1
				} else {
					publicip.Status = 0
				}
				tmpnodemap.Ippublic[id] = publicip
				g.NodeMap[NetId][NodeId] = tmpnodemap
				//可以先检测再计算下发
				g.CompRout(NetId, NodeId) //计算出cfg
				g.InitCfg(NetId)
				break
			}
		}

		for id, privateip := range tmpnodemap.Ipprivate {
			if privateip.Ip == ip {
				if types == "start" {
					privateip.Status = 1
				} else {
					privateip.Status = 0
				}
				tmpnodemap.Ipprivate[id] = privateip
				g.NodeMap[NetId][NodeId] = tmpnodemap
				g.CompRout(NetId, NodeId) //计算出cfg
				g.InitCfg(NetId)
				break
			}
		}

		response.Status = 1
		response.Data = "success"
		this.Data["json"] = response
		this.ServeJSON()

	} else {
		fmt.Println("Key Not Found")
		response.Status = 0
		response.Data = "not fund " + NodeId
		this.Data["json"] = response
		this.ServeJSON()
	}

}
