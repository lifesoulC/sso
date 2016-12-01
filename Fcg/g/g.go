package g

import "fmt"

//import "net"

var NodeMap map[string]map[string]Nodelinkmap
var RouteMap map[string]map[string]map[string]RoutlinkMap
var RouteMapNet map[string]map[string]map[string]RoutelinkMapNet
var PollMapNet map[string]map[string][]byte
var PollMapNode map[string]map[string][]byte

var LbRouteNode map[string]string
var LbRouteNet map[string]string
var LbRoutetmp []LbRoute

//var Nodemapinfo Nodemap

func InitEnv() {
	//Nodemapinfo := Nodemap{}
	//fmt.Println(Nodemapinfo)
	NodeMap = make(map[string]map[string]Nodelinkmap)
	RouteMap = make(map[string]map[string]map[string]RoutlinkMap)        //net,node,dst RoutlinkMap
	RouteMapNet = make(map[string]map[string]map[string]RoutelinkMapNet) //net,node,dst RoutelinkMapNet
	PollMapNet = make(map[string]map[string][]byte)
	PollMapNode = make(map[string]map[string][]byte)

	LbRouteNode = make(map[string]string)
	LbRouteNet = make(map[string]string)
	LbRoutetmp = []LbRoute{}

	//	tmpnode := Nodelinkmap{}
	//	tmpnode.Flow = "1"
	//	tmpnode.HostType = "vm"
	//	tmpnode.Id = "1-1"
	//	tmpnode.Ip = "123.123.123.128"

	//	tmpprivate := IPinfo{}
	//	tmpprivate.Ip = "10.10.1.1"
	//	tmpprivate.Isp = ""
	//	tmpprivate.Status = 1
	//	tmpnode.Ipprivate = append(tmpnode.Ipprivate, tmpprivate)

	//	tmppublic := IPinfo{}
	//	tmppublic.Ip = "123.123.123.128"
	//	tmppublic.Isp = "电信"
	//	tmppublic.Status = 1
	//	tmpnode.Ippublic = append(tmpnode.Ippublic, tmppublic)

	//	tmpnode.Location = "北京"
	//	tmpnode.Name = "bjvm"
	//	tmpnode.Net = "1"
	//	tmpnode.Node = "1"
	//	tmpnode.Status = 1

	//	tmpnodemap := make(map[string]Nodelinkmap)
	//	tmpnodemap["1"] = tmpnode

	//	////////////////////////////////
	//	tmpnode1 := Nodelinkmap{}
	//	tmpnode1.Flow = "1"
	//	tmpnode1.HostType = "vm"
	//	tmpnode1.Id = "1-3"
	//	tmpnode1.Ip = "123.123.123.124"

	//	tmpprivate1 := IPinfo{}
	//	tmpprivate1.Ip = "10.10.1.2"
	//	tmpprivate1.Isp = ""
	//	tmpprivate1.Status = 1
	//	tmpnode1.Ipprivate = append(tmpnode1.Ipprivate, tmpprivate)

	//	tmppublic1 := IPinfo{}
	//	tmppublic1.Ip = "123.123.123.124"
	//	tmppublic1.Isp = "电信"
	//	tmppublic1.Status = 1
	//	tmpnode1.Ippublic = append(tmpnode1.Ippublic, tmppublic1)

	//	tmpnode1.Location = "上海"
	//	tmpnode1.Name = "shvm"
	//	tmpnode1.Net = "1"
	//	tmpnode1.Node = "3"
	//	tmpnode1.Status = 1

	//	//tmpnodemap := make(map[string]Nodelinkmap)
	//	tmpnodemap["3"] = tmpnode1

	//	LbRouteNode["1-3"] = `NODE ROUTE: (dst_node) 1-1 (flow_level) 1 (cur_node) 1-3 (next_node) 1-1 (cur_ip) 61.172.243.106 (next_ip) 118.187.3.6
	//	NODE ROUTE: (dst_node) 1-1 (flow_level) 2 (cur_node) 1-3 (next_node) 1-1 (cur_ip) 61.172.243.106 (next_ip) 118.187.3.6
	//	NODE ROUTE: (dst_node) 1-2 (flow_level) 1 (cur_node) 1-3 (next_node) 1-4 (cur_ip) 61.172.243.106 (next_ip) 139.224.19.178
	//	NODE ROUTE: (dst_node) 1-2 (flow_level) 2 (cur_node) 1-3 (next_node) 1-4 (cur_ip) 61.172.243.106 (next_ip) 139.224.19.178
	//	NODE ROUTE: (dst_node) 1-4 (flow_level) 1 (cur_node) 1-3 (next_node) 1-4 (cur_ip) 61.172.243.106 (next_ip) 139.224.19.178
	//	NODE ROUTE: (dst_node) 1-4 (flow_level) 2 (cur_node) 1-3 (next_node) 1-4 (cur_ip) 61.172.243.106 (next_ip) 139.224.19.178
	//	NODE ROUTE: (dst_node) 1-5 (flow_level) 1 (cur_node) 1-3 (next_node) 1-5 (cur_ip) 61.172.243.106 (next_ip) 183.60.189.202
	//	NODE ROUTE: (dst_node) 1-5 (flow_level) 2 (cur_node) 1-3 (next_node) 1-5 (cur_ip) 61.172.243.106 (next_ip) 183.60.189.202
	//	`

	//	LbRouteNet["1-1"] = `NET ROUTE: (dst_net) 2 (flow_level) 1 (cur_node) 1-3 (next_node) 1-5 (cur_ip) 61.172.243.106 (next_ip) 183.60.189.202
	//	NET ROUTE: (dst_net) 2 (flow_level) 2 (cur_node) 1-3 (next_node) 1-1 (cur_ip) 61.172.243.106 (next_ip) 118.187.3.6
	//	NET ROUTE: (dst_net) 3 (flow_level) 1 (cur_node) 1-3 (next_node) 1-5 (cur_ip) 61.172.243.106 (next_ip) 183.60.189.202
	//	NET ROUTE: (dst_net) 3 (flow_level) 2 (cur_node) 1-3 (next_node) 3-1 (cur_ip) 10.128.0.129 (next_ip) 10.128.0.137
	//	NET ROUTE: (dst_net) 4 (flow_level) 1 (cur_node) 1-3 (next_node) 1-5 (cur_ip) 61.172.243.106 (next_ip) 183.60.189.202
	//	NET ROUTE: (dst_net) 4 (flow_level) 2 (cur_node) 1-3 (next_node) 1-5 (cur_ip) 61.172.243.106 (next_ip) 183.60.189.202`
	//	///////////////////////////

	//	tmpnode1 = Nodelinkmap{}
	//	tmpnode1.Flow = "1"
	//	tmpnode1.HostType = "vm"
	//	tmpnode1.Id = "1-5"
	//	tmpnode1.Ip = "123.123.123.125"

	//	tmpprivate1 = IPinfo{}
	//	tmpprivate1.Ip = "10.10.1.3"
	//	tmpprivate1.Isp = ""
	//	tmpprivate1.Status = 1
	//	tmpnode1.Ipprivate = append(tmpnode1.Ipprivate, tmpprivate1)

	//	tmppublic1 = IPinfo{}
	//	tmppublic1.Ip = "123.123.123.125"
	//	tmppublic1.Isp = "电信"
	//	tmppublic1.Status = 1
	//	tmpnode1.Ippublic = append(tmpnode1.Ippublic, tmppublic1)

	//	tmpnode1.Location = "广州"
	//	tmpnode1.Name = "gzvm"
	//	tmpnode1.Net = "1"
	//	tmpnode1.Node = "4"
	//	tmpnode1.Status = 1

	//	tmpnodemap["5"] = tmpnode1
	//	NodeMap["1"] = tmpnodemap

	//	LbRouteNet["1-5"] = `NET ROUTE: (dst_net) 2 (flow_level) 1 (cur_node) 1-5 (next_node) 4-1 (cur_ip) 206.161.224.130 (next_ip) 119.81.234.186
	//	NET ROUTE: (dst_net) 2 (flow_level) 2 (cur_node) 1-5 (next_node) 1-1 (cur_ip) 183.60.189.202 (next_ip) 118.187.3.6
	//	NET ROUTE: (dst_net) 3 (flow_level) 1 (cur_node) 1-5 (next_node) 4-1 (cur_ip) 206.161.224.130 (next_ip) 119.81.234.186
	//	NET ROUTE: (dst_net) 3 (flow_level) 2 (cur_node) 1-5 (next_node) 1-3 (cur_ip) 183.60.189.202 (next_ip) 61.172.243.106
	//	NET ROUTE: (dst_net) 4 (flow_level) 1 (cur_node) 1-5 (next_node) 4-1 (cur_ip) 206.161.224.130 (next_ip) 119.81.234.186
	//	NET ROUTE: (dst_net) 4 (flow_level) 2 (cur_node) 1-5 (next_node) 4-1 (cur_ip) 206.161.224.130 (next_ip) 119.81.234.18`

	//	LbRouteNode["1-5"] = `NODE ROUTE: (dst_node) 1-1 (flow_level) 1 (cur_node) 1-5 (next_node) 1-1 (cur_ip) 183.60.189.202 (next_ip) 118.187.3.6
	//	NODE ROUTE: (dst_node) 1-1 (flow_level) 2 (cur_node) 1-5 (next_node) 1-1 (cur_ip) 183.60.189.202 (next_ip) 118.187.3.6
	//	NODE ROUTE: (dst_node) 1-2 (flow_level) 1 (cur_node) 1-5 (next_node) 1-2 (cur_ip) 183.60.189.202 (next_ip) 123.57.6.17
	//	NODE ROUTE: (dst_node) 1-2 (flow_level) 2 (cur_node) 1-5 (next_node) 1-2 (cur_ip) 183.60.189.202 (next_ip) 123.57.6.17
	//	NODE ROUTE: (dst_node) 1-3 (flow_level) 1 (cur_node) 1-5 (next_node) 1-3 (cur_ip) 183.60.189.202 (next_ip) 61.172.243.106
	//	NODE ROUTE: (dst_node) 1-3 (flow_level) 2 (cur_node) 1-5 (next_node) 1-3 (cur_ip) 183.60.189.202 (next_ip) 61.172.243.106
	//	NODE ROUTE: (dst_node) 1-4 (flow_level) 1 (cur_node) 1-5 (next_node) 1-4 (cur_ip) 183.60.189.202 (next_ip) 139.224.19.178
	//	NODE ROUTE: (dst_node) 1-4 (flow_level) 2 (cur_node) 1-5 (next_node) 1-4 (cur_ip) 183.60.189.202 (next_ip) 139.224.19.178
	//	`
	//	////////////////////////////////////////
	//	tmpnode1 = Nodelinkmap{}
	//	tmpnode1.Flow = "1"
	//	tmpnode1.HostType = "vm"
	//	tmpnode1.Id = "3-1"
	//	tmpnode1.Ip = "123.123.123.126"

	//	tmpprivate1 = IPinfo{}
	//	tmpprivate1.Ip = "10.10.1.4"
	//	tmpprivate1.Isp = ""
	//	tmpprivate1.Status = 1
	//	tmpnode1.Ipprivate = append(tmpnode1.Ipprivate, tmpprivate1)

	//	tmppublic1 = IPinfo{}
	//	tmppublic1.Ip = "123.123.123.126"
	//	tmppublic1.Isp = ""
	//	tmppublic1.Status = 1
	//	tmpnode1.Ippublic = append(tmpnode1.Ippublic, tmppublic1)

	//	tmpnode1.Location = "华盛顿"
	//	tmpnode1.Name = "us03"
	//	tmpnode1.Net = "3"
	//	tmpnode1.Node = "1"
	//	tmpnode1.Status = 1
	//	tmpnodemap1 := make(map[string]Nodelinkmap)
	//	tmpnodemap1["1"] = tmpnode1
	//	NodeMap["2"] = tmpnodemap1
	//	LbRouteNet["3-1"] = `"NET ROUTE: (dst_net) 1 (flow_level) 1 (cur_node) 3-1 (next_node) 3-2 (cur_ip) 198.148.124.38 (next_ip) 50.23.125.226
	//	                         NET ROUTE: (dst_net) 1 (flow_level) 2 (cur_node) 3-1 (next_node) 1-3 (cur_ip) 10.128.0.137 (next_ip) 10.128.0.129
	//	                         NET ROUTE: (dst_net) 2 (flow_level) 1 (cur_node) 3-1 (next_node) 3-2 (cur_ip) 198.148.124.38 (next_ip) 50.23.125.226
	//	                         NET ROUTE: (dst_net) 2 (flow_level) 2 (cur_node) 3-1 (next_node) 3-2 (cur_ip) 198.148.124.38 (next_ip) 50.23.125.226
	//	                         NET ROUTE: (dst_net) 4 (flow_level) 1 (cur_node) 3-1 (next_node) 3-2 (cur_ip) 198.148.124.38 (next_ip) 50.23.125.226
	//	                         NET ROUTE: (dst_net) 4 (flow_level) 2 (cur_node) 3-1 (next_node) 3-2 (cur_ip) 198.148.124.38 (next_ip) 50.23.125.226
	//	                         "`
	//	LbRouteNode["3-1"] = `NODE ROUTE: (dst_node) 3-2 (flow_level) 1 (cur_node) 3-1 (next_node) 3-2 (cur_ip) 198.148.124.38 (next_ip) 50.23.125.226
	//	NODE ROUTE: (dst_node) 3-2 (flow_level) 2 (cur_node) 3-1 (next_node) 3-2 (cur_ip) 198.148.124.38 (next_ip) 50.23.125.226
	//	`
	//	///////////////////////
	//	tmpnode1 = Nodelinkmap{}
	//	tmpnode1.Flow = "1"
	//	tmpnode1.HostType = "sl"
	//	tmpnode1.Id = "4-1"
	//	tmpnode1.Ip = "123.123.123.127"

	//	tmpprivate1 = IPinfo{}
	//	tmpprivate1.Ip = "10.10.1.5"
	//	tmpprivate1.Isp = ""
	//	tmpprivate1.Status = 1
	//	tmpnode1.Ipprivate = append(tmpnode1.Ipprivate, tmpprivate1)

	//	tmppublic1 = IPinfo{}
	//	tmppublic1.Ip = "123.123.123.127"
	//	tmppublic1.Isp = "电信"
	//	tmppublic1.Status = 1
	//	tmpnode1.Ippublic = append(tmpnode1.Ippublic, tmppublic1)

	//	tmpnode1.Location = "香港"
	//	tmpnode1.Name = "hk"
	//	tmpnode1.Net = "4"
	//	tmpnode1.Node = "1"
	//	tmpnode1.Status = 1
	//	tmpnodemap2 := make(map[string]Nodelinkmap)
	//	tmpnodemap2["1"] = tmpnode1
	//	NodeMap["4"] = tmpnodemap2
	//////////////////////
	fmt.Println(NodeMap)
	fmt.Println(RouteMap)
	fmt.Println(RouteMapNet)
}
