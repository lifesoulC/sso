package controllers

import (
	"Fcg/g"
	"encoding/json"
	"fmt"
	//"os"
	"strconv"
	"strings"
)

func (this *NodeController) Path() {
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
	response.Data.Net = nil
	response.Data.Nodes = nil

	for _, nodemap := range g.RouteMap {
		for _, nodemap := range nodemap {
			for _, routemap := range nodemap {
				response.Data.Nodes = append(response.Data.Nodes, routemap)
			}
		}
	}

	for _, nodemap := range g.RouteMapNet {
		for _, nodemap := range nodemap {
			for _, routemap := range nodemap {
				response.Data.Net = append(response.Data.Net, routemap)
			}
		}
	}
	//fmt.Println("respons:", response)
	//fmt.Println("respons.data.nodes:", response.Data.Nodes)
	//fmt.Println("respons.data.net:", response.Data.Net)
	if response.Data.Nodes == nil {
		response.Status = 0
		//fmt.Println("Status = 0")
		this.Data["json"] = response
		this.ServeJSON()
	} else {
		response.Status = 1
		//fmt.Println("Status = 1")
		this.Data["json"] = response
		this.ServeJSON()

	}
}

func (this *NodeController) ChangePath() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	var err error
	type Node struct {
		NodeId   string `json:"nodeId"`
		Ip       string `json:"ip"`
		Isp      string `json:"isp"`
		Location string `json:"location"`
	}
	type Data struct {
		Types string `json:"types"`
		Level int    `json:"level"`
		From  Node   `json:"from"`
		To    Node   `json:"to"`
	}
	type Response struct {
		Status int    `json:"status"`
		Data   string `json:"datas"`
	}
	response := Response{}
	datatmp := Data{}
	datatmp.From.NodeId = this.GetString("from[nodeId]")
	datatmp.From.Ip = this.GetString("from[ip]")
	datatmp.From.Isp = this.GetString("from[isp]")
	datatmp.From.Location = this.GetString("from[location]")

	datatmp.To.NodeId = this.GetString("to[nodeId]")
	datatmp.To.Ip = this.GetString("to[ip]")
	datatmp.To.Isp = this.GetString("to[isp]")
	datatmp.To.Location = this.GetString("to[location]")

	//	datatmp.Level, err = this.GetInt("level")
	//	if err != nil {
	//		response.Status = 0
	//		response.Data = "not fund this addr"
	//		this.Data["json"] = response
	//		this.ServeJSON()
	//	}

	strlevel := this.GetString("level")
	b, err := strconv.Atoi(strlevel)
	if err != nil {
		fmt.Println("字符串转换成整数失败")
	}
	datatmp.Level = b
	datatmp.Types = this.GetString("type")

	netstr := strings.Split(datatmp.From.NodeId, "-") //解析发来的字符
	//fmt.Println(netstr)
	srcNetId := netstr[0]
	srcNodeId := netstr[1]

	nodestr := strings.Split(datatmp.To.NodeId, "-") //解析发来的字符
	dstNetId := nodestr[0]
	dstNodeId := nodestr[1]

	if datatmp.Types == "add" {
		if srcNetId == dstNetId {
			routelinkmap := g.RoutlinkMap{}

			routelinkmap.From.NodeId = datatmp.From.NodeId
			routelinkmap.From.Location = datatmp.From.Location
			routelinkmap.From.Ip = datatmp.From.Ip
			routelinkmap.From.Isp = datatmp.From.Isp

			routelinkmap.To.NodeId = datatmp.To.NodeId
			routelinkmap.To.Location = datatmp.To.Location
			routelinkmap.To.Ip = datatmp.To.Ip
			routelinkmap.To.Isp = datatmp.To.Isp

			routelinkmap.Level = datatmp.Level
			routelinkmap.Delays = 300
			routelinkmap.Lost = ""
			routelinkmap.Status = 1

			if _, ok := g.RouteMap[srcNetId]; ok {
				if _, ok := g.RouteMap[srcNetId][srcNodeId]; ok {
					//					if _, ok := g.RouteMap[srcNetId][srcNodeId][datatmp.To.NodeId]; ok {
					//						g.RouteMap[srcNetId][srcNodeId][datatmp.To.NodeId] = routelinkmap
					//					} else {
					g.RouteMap[srcNetId][srcNodeId][datatmp.To.NodeId] = routelinkmap

					err := ChangeCfg(srcNetId, srcNodeId, "node")
					if err != nil {
						response.Status = 0
						response.Data = "cfg 下发失败"
						this.Data["json"] = response
						this.ServeJSON()
					}
				} else {
					dstmap := make(map[string]g.RoutlinkMap)
					dstmap[datatmp.To.NodeId] = routelinkmap
					g.RouteMap[srcNetId][srcNodeId] = dstmap
					//生成配置下发
					err := ChangeCfg(srcNetId, srcNodeId, "node")
					if err != nil {
						response.Status = 0
						response.Data = "cfg 下发失败"
						this.Data["json"] = response
						this.ServeJSON()
					}
				}
			} else {
				dstmap := make(map[string]g.RoutlinkMap)
				dstmap[datatmp.To.NodeId] = routelinkmap
				nodemap := make(map[string]map[string]g.RoutlinkMap)
				nodemap[srcNodeId] = dstmap
				g.RouteMap[srcNetId] = nodemap
				//	生成 cfg 下发
				err := ChangeCfg(srcNetId, srcNodeId, "node")
				if err != nil {
					response.Status = 0
					response.Data = "cfg 下发失败"
					this.Data["json"] = response
					this.ServeJSON()
				}
			}

			routelinkmap.From.NodeId = datatmp.To.NodeId
			routelinkmap.From.Location = datatmp.To.Location
			routelinkmap.From.Ip = datatmp.To.Ip
			routelinkmap.From.Isp = datatmp.To.Isp

			routelinkmap.To.NodeId = datatmp.From.NodeId
			routelinkmap.To.Location = datatmp.From.Location
			routelinkmap.To.Ip = datatmp.From.Ip
			routelinkmap.To.Isp = datatmp.From.Isp

			routelinkmap.Level = datatmp.Level
			routelinkmap.Delays = 300
			routelinkmap.Lost = ""
			routelinkmap.Status = 1

			if _, ok := g.RouteMap[dstNetId]; ok {
				if _, ok := g.RouteMap[dstNetId][dstNodeId]; ok {
					//					if _, ok := g.RouteMap[srcNetId][srcNodeId][datatmp.To.NodeId]; ok {
					//						g.RouteMap[srcNetId][srcNodeId][datatmp.To.NodeId] = routelinkmap
					//					} else {
					g.RouteMap[dstNetId][dstNodeId][datatmp.From.NodeId] = routelinkmap

					err := ChangeCfg(dstNetId, dstNodeId, "node")
					if err != nil {
						response.Status = 0
						response.Data = "cfg 下发失败"
						this.Data["json"] = response
						this.ServeJSON()
					}
				} else {
					dstmap := make(map[string]g.RoutlinkMap)
					dstmap[datatmp.From.NodeId] = routelinkmap
					g.RouteMap[dstNetId][dstNodeId] = dstmap
					//生成配置下发
					err := ChangeCfg(dstNetId, dstNodeId, "node")
					if err != nil {
						response.Status = 0
						response.Data = "cfg 下发失败"
						this.Data["json"] = response
						this.ServeJSON()
					}
				}
			} else {
				dstmap := make(map[string]g.RoutlinkMap)
				dstmap[datatmp.From.NodeId] = routelinkmap
				nodemap := make(map[string]map[string]g.RoutlinkMap)
				nodemap[dstNodeId] = dstmap
				g.RouteMap[dstNetId] = nodemap
				//	生成 cfg 下发
				err := ChangeCfg(dstNetId, dstNodeId, "node")
				if err != nil {
					response.Status = 0
					response.Data = "cfg 下发失败"
					this.Data["json"] = response
					this.ServeJSON()
				}
			}
			//生成 cfg 下发

			response.Status = 1
			response.Data = "success"
			this.Data["json"] = response
			this.ServeJSON()
		}
		if srcNetId != dstNetId {
			routelinkmap := g.RoutelinkMapNet{}

			routelinkmap.From.NodeId = datatmp.From.NodeId
			routelinkmap.From.Location = datatmp.From.Location
			routelinkmap.From.Ip = datatmp.From.Ip
			routelinkmap.From.Isp = datatmp.From.Isp

			routelinkmap.To.NodeId = datatmp.To.NodeId
			routelinkmap.To.Location = datatmp.To.Location
			routelinkmap.To.Ip = datatmp.To.Ip
			routelinkmap.To.Isp = datatmp.To.Isp

			routelinkmap.Level = datatmp.Level
			routelinkmap.Delays = 300
			routelinkmap.Lost = ""
			routelinkmap.Status = 1

			if _, ok := g.RouteMapNet[srcNetId]; ok {
				if _, ok := g.RouteMapNet[srcNetId][srcNodeId]; ok {
					//					if _, ok := g.RouteMap[srcNetId][srcNodeId][datatmp.To.NodeId]; ok {
					//						g.RouteMap[srcNetId][srcNodeId][datatmp.To.NodeId] = routelinkmap
					//					} else {
					g.RouteMapNet[srcNetId][srcNodeId][datatmp.To.NodeId] = routelinkmap

					err := ChangeCfg(srcNetId, srcNodeId, "net")
					if err != nil {
						response.Status = 0
						response.Data = "cfg 下发失败"
						this.Data["json"] = response
						this.ServeJSON()
					}
				} else {
					dstmap := make(map[string]g.RoutelinkMapNet)
					dstmap[datatmp.To.NodeId] = routelinkmap
					g.RouteMapNet[srcNetId][srcNodeId] = dstmap
					//生成配置下发
					err := ChangeCfg(srcNetId, srcNodeId, "net")
					if err != nil {
						response.Status = 0
						response.Data = "cfg 下发失败"
						this.Data["json"] = response
						this.ServeJSON()
					}
				}
			} else {
				dstmap := make(map[string]g.RoutelinkMapNet)
				dstmap[datatmp.To.NodeId] = routelinkmap
				nodemap := make(map[string]map[string]g.RoutelinkMapNet)
				nodemap[srcNodeId] = dstmap
				g.RouteMapNet[srcNetId] = nodemap
				//	生成 cfg 下发
				err := ChangeCfg(srcNetId, srcNodeId, "net")
				if err != nil {
					response.Status = 0
					response.Data = "cfg 下发失败"
					this.Data["json"] = response
					this.ServeJSON()
				}
			}

			routelinkmap.From.NodeId = datatmp.To.NodeId
			routelinkmap.From.Location = datatmp.To.Location
			routelinkmap.From.Ip = datatmp.To.Ip
			routelinkmap.From.Isp = datatmp.To.Isp

			routelinkmap.To.NodeId = datatmp.From.NodeId
			routelinkmap.To.Location = datatmp.From.Location
			routelinkmap.To.Ip = datatmp.From.Ip
			routelinkmap.To.Isp = datatmp.From.Isp

			routelinkmap.Level = datatmp.Level
			routelinkmap.Delays = 300
			routelinkmap.Lost = ""
			routelinkmap.Status = 1

			if _, ok := g.RouteMapNet[dstNetId]; ok {
				if _, ok := g.RouteMapNet[dstNetId][dstNodeId]; ok {
					//					if _, ok := g.RouteMap[srcNetId][srcNodeId][datatmp.To.NodeId]; ok {
					//						g.RouteMap[srcNetId][srcNodeId][datatmp.To.NodeId] = routelinkmap
					//					} else {
					g.RouteMapNet[dstNetId][dstNodeId][datatmp.From.NodeId] = routelinkmap
					//					}
					//生成配置下发
					err := ChangeCfg(dstNetId, dstNodeId, "net")
					if err != nil {
						response.Status = 0
						response.Data = "cfg 下发失败"
						this.Data["json"] = response
						this.ServeJSON()
					}
				} else {
					dstmap := make(map[string]g.RoutelinkMapNet)
					dstmap[datatmp.From.NodeId] = routelinkmap
					g.RouteMapNet[dstNetId][dstNodeId] = dstmap
					//生成配置下发
					err := ChangeCfg(dstNetId, dstNodeId, "net")
					if err != nil {
						response.Status = 0
						response.Data = "cfg 下发失败"
						this.Data["json"] = response
						this.ServeJSON()
					}
				}
			} else {
				dstmap := make(map[string]g.RoutelinkMapNet)
				dstmap[datatmp.From.NodeId] = routelinkmap
				nodemap := make(map[string]map[string]g.RoutelinkMapNet)
				nodemap[dstNodeId] = dstmap
				g.RouteMapNet[dstNetId] = nodemap
				//	生成 cfg 下发
				err := ChangeCfg(dstNetId, dstNodeId, "net")
				if err != nil {
					response.Status = 0
					response.Data = "cfg 下发失败"
					this.Data["json"] = response
					this.ServeJSON()
				}
			}
			//	生成 cfg 下发

			response.Status = 1
			response.Data = "success"
			this.Data["json"] = response
			this.ServeJSON()

		}
	} //var RouteMapNet map[string]map[string]map[string]RoutelinkMapNet
	if datatmp.Types == "del" {
		if srcNetId != dstNetId {
			if _, ok := g.RouteMapNet[srcNetId][srcNodeId][datatmp.To.NodeId]; ok {
				if _, ok := g.RouteMapNet[dstNetId][dstNodeId][datatmp.From.NodeId]; ok {
					delete(g.RouteMapNet[srcNetId][srcNodeId], datatmp.To.NodeId)
					delete(g.RouteMapNet[dstNetId][dstNodeId], datatmp.From.NodeId)

					response.Status = 1
					response.Data = "success"
					this.Data["json"] = response
					this.ServeJSON()
					//生成配置 下发

					err := ChangeCfg(srcNetId, srcNodeId, "net")
					if err != nil {
						response.Status = 0
						response.Data = "not fund this addr"
						this.Data["json"] = response
						this.ServeJSON()
					}

					err = ChangeCfg(dstNetId, dstNodeId, "net")
					if err != nil {
						response.Status = 0
						response.Data = "not fund this addr"
						this.Data["json"] = response
						this.ServeJSON()
					}

				} else {
					//返回错误信息
					fmt.Println("404 error")
					response.Status = 0
					response.Data = "not fund this addr"
					this.Data["json"] = response
					this.ServeJSON()
				}
			} else {
				//返回错误信息
				fmt.Println("411 error")
				response.Status = 0
				response.Data = "not fund this addr"
				this.Data["json"] = response
				this.ServeJSON()
			}
		} else {
			if _, ok := g.RouteMap[srcNetId][srcNodeId][datatmp.To.NodeId]; ok {
				if _, ok := g.RouteMap[dstNetId][dstNodeId][datatmp.From.NodeId]; ok {
					delete(g.RouteMap[srcNetId][srcNodeId], datatmp.To.NodeId)
					delete(g.RouteMap[dstNetId][dstNodeId], datatmp.From.NodeId)
					//生成配置 下发
					err := ChangeCfg(srcNetId, srcNodeId, "node")
					if err != nil {
						response.Status = 0
						response.Data = "not fund this addr"
						this.Data["json"] = response
						this.ServeJSON()
					}

					err = ChangeCfg(dstNetId, dstNodeId, "node")
					if err != nil {
						response.Status = 0
						response.Data = "not fund this addr"
						this.Data["json"] = response
						this.ServeJSON()
					}

					response.Status = 1
					fmt.Println("439 error")
					response.Data = "success"
					this.Data["json"] = response
					this.ServeJSON()
				} else {
					//返回错误信息
					fmt.Println("445 error")
					response.Status = 0
					response.Data = "not fund this addr"
					this.Data["json"] = response
					this.ServeJSON()
				}
			} else {
				//返回错误信息
				fmt.Println("453 error")
				response.Status = 0
				response.Data = "not fund this addr"
				this.Data["json"] = response
				this.ServeJSON()
			}
		}
	}

	//fmt.Println(g.RouteMap)

	//fmt.Println("RoutNEt:", g.RouteMapNet)
}

//var RouteMap map[string]map[string]map[string]RoutlinkMap
//var RouteMapNet map[string]map[string]map[string]RoutelinkMapNet
//var NodeMap map[string]map[string]Nodelinkmap
func ChangeCfg(net string, node string, Types string) error {
	respcfgs := []g.RespCfg{}
	if Types == "node" {
		for _, nodemapinfo := range g.RouteMap[net][node] {
			respcfg := g.RespCfg{}
			respcfg.SrcAddr = nodemapinfo.From.NodeId
			respcfg.SrcIPv4 = nodemapinfo.From.Ip

			respcfg.DstAddr = nodemapinfo.To.NodeId
			respcfg.DstIPv4 = nodemapinfo.To.Ip

			respcfg.FlowLevel = 1

			respcfgs = append(respcfgs, respcfg)
		}
		b, _ := json.MarshalIndent(respcfgs, "", "")
		fmt.Println("path 477:", string(b))
		if len(respcfgs) == 0 {
			str := "[]"
			b = []byte(str)
			if _, ok := g.PollMapNode[net]; ok {
				if _, ok := g.PollMapNode[net][node]; ok {
					g.PollMapNode[net][node] = b
				}
			} else {
				dstmap := make(map[string][]byte)
				dstmap[node] = b
				g.PollMapNode[net] = dstmap
			}
		} else {
			fmt.Println("520 path respcfgs", respcfgs)
			b, _ = json.MarshalIndent(respcfgs, "", "")
			if _, ok := g.PollMapNode[net]; ok {
				//if _, ok := g.PollMapNet[net][node]; ok {
				g.PollMapNode[net][node] = b
				//}
			} else {
				dstmap := make(map[string][]byte)
				dstmap[node] = b
				g.PollMapNode[net] = dstmap
				fmt.Println("path 530 g.PollMapNode", g.PollMapNode[net])
			}
		}

		//		ip := g.NodeMap[net][node].Ippublic[0].Ip
		//		path := "http://" + ip + ":9900/donodelink"
		//		reqbody, err := g.HttpPostCfg(b, path) //下发
		//		if (err != nil) || (reqbody == "") {
		//			fmt.Println(err)
		//			return err
		//		}
		//		strbody := string(reqbody)
		//		if strbody == "OK" {
		//			fmt.Println(net + "-" + node + "success")
		//			return nil
		//		}
	} else {
		for _, nodemapinfo := range g.RouteMapNet[net][node] {
			respcfg := g.RespCfg{}
			respcfg.SrcAddr = nodemapinfo.From.NodeId
			respcfg.SrcIPv4 = nodemapinfo.From.Ip

			respcfg.DstAddr = nodemapinfo.To.NodeId
			respcfg.DstIPv4 = nodemapinfo.To.Ip

			respcfg.FlowLevel = 1

			respcfgs = append(respcfgs, respcfg)
		}
		b, _ := json.MarshalIndent(respcfgs, "", "")
		fmt.Println("path 505", string(b))
		if len(respcfgs) == 0 {
			str := "[]"
			b = []byte(str)
			if _, ok := g.PollMapNet[net]; ok {
				if _, ok := g.PollMapNet[net][node]; ok {
					g.PollMapNet[net][node] = b
				}
			} else {
				dstmap := make(map[string][]byte)
				dstmap[node] = b
				g.PollMapNet[net] = dstmap
			}
		} else {
			fmt.Println("520 path respcfgs", respcfgs)
			b, _ = json.MarshalIndent(respcfgs, "", "")
			if _, ok := g.PollMapNet[net]; ok {
				//if _, ok := g.PollMapNet[net][node]; ok {
				g.PollMapNet[net][node] = b
				//}
			} else {
				dstmap := make(map[string][]byte)
				dstmap[node] = b
				g.PollMapNet[net] = dstmap
				fmt.Println("path 530 g.PollMapNet", g.PollMapNet[net])
			}
		}
		//		ip := g.NodeMap[net][node].Ippublic[0].Ip
		//		path := "http://" + ip + ":9900/donetlink"
		//		reqbody, err := g.HttpPostCfg(b, path) //下发
		//		if (err != nil) || (reqbody == "") {
		//			fmt.Println(err)
		//			return err
		//		}
		//		strbody := string(reqbody)
		//		if strbody == "OK" {
		//			fmt.Println(net + "-" + node + "success")
		//			return nil
		//		}
	}
	//fmt.Println("hallo respcfgs:", respcfgs)

	return nil
}
