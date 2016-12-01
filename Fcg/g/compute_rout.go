package g

import (
	"fmt"
	"strings"
)

func RoutePath(srcaddr string, dstaddr string, leve string) { //"1-1" "1-3" "2"
	srcStr := LbRouteNode[srcaddr]
	fmt.Println("srcStr", srcStr)
	if srcStr == "" {
		return
	}

	//dstaddr := RoutNetMap[dstaddr]
	lbRoutetmp := LbRoute{}
	Lstrs := strings.SplitN(srcStr, "\n", -1) //得出每行的信息 NODE ROUTE: (dst_node) 1-2 (flow_level) 1 (cur_node) 1-1 (next_node) 1-2 (cur_ip) 61.172.243.105 (next_ip) 183.129.179.162
	for _, v := range Lstrs {
		if v == "" {
			break
		}
		Wstrs := strings.SplitN(v, " ", -1) //得出每个单词的信息 "NODE" "ROUTE".....

		if (Wstrs[3] == dstaddr) && (Wstrs[5] == leve) && (Wstrs[9] == dstaddr) {

			lbRoutetmp.DstAddr = Wstrs[9]
			lbRoutetmp.Dstip = Wstrs[13]

			lbRoutetmp.SrcAddr = Wstrs[7]
			lbRoutetmp.SrcIp = Wstrs[11]

			LbRoutetmp = append(LbRoutetmp, lbRoutetmp)

			return
		} else {
			if (Wstrs[3] == dstaddr) && (Wstrs[5] == leve) {
				//info := NetNameMap[Wstrs[9]]
				//Route = Route + "(" + Wstrs[11] + ")" + "->" + info.Name + "(" + Wstrs[13] + ")"
				lbRoutetmp.DstAddr = Wstrs[9]
				lbRoutetmp.Dstip = Wstrs[13]

				lbRoutetmp.SrcAddr = Wstrs[7]
				lbRoutetmp.SrcIp = Wstrs[11]

				LbRoutetmp = append(LbRoutetmp, lbRoutetmp)

				RoutePath(Wstrs[9], dstaddr, leve)
			}
		}
	}

}

func RouteNetPath(srcaddr string, dstaddr string, leve string) { //"1-1" "1-3" "2"
	dstaddrs := strings.Split(dstaddr, "-") //将"3-1" 解析出来
	srcStr := LbRouteNet[srcaddr]

	lbRoutetmp := LbRoute{}
	if srcStr == "" {
		//	Route = "get message error"
		return
	}
	Lstrs := strings.SplitN(srcStr, "\n", -1)
	for _, v := range Lstrs {
		if v == "" {
			break
		}
		Wstrs := strings.SplitN(v, " ", -1) //得出每个单词的信息 "NODE" "ROUTE".....
		if (Wstrs[3] == dstaddrs[0]) && (Wstrs[5] == leve) {
			//			info := NetNameMap[Wstrs[9]]
			//			Route = Route + "(" + Wstrs[11] + ")" + "->" + info.Name + "(" + Wstrs[13] + ")"
			lbRoutetmp.DstAddr = Wstrs[9]
			lbRoutetmp.Dstip = Wstrs[13]

			lbRoutetmp.SrcAddr = Wstrs[7]
			lbRoutetmp.SrcIp = Wstrs[11]

			LbRoutetmp = append(LbRoutetmp, lbRoutetmp)
			nextnode := strings.Split(Wstrs[9], "-")

			if Wstrs[9] == dstaddr {
				return
			}
			if nextnode[0] != dstaddrs[0] { //如果本机不是出口机器
				RouteNetPath(Wstrs[9], dstaddr, leve)
			} else { //本机是出口机器
				RoutePath(Wstrs[9], dstaddr, leve)
			}
		}

	}
}

func Routeing(srcaddr string, dstaddr string, leve string) {
	dstaddrs := strings.Split(dstaddr, "-") //解析发来的字符
	srcaddrs := strings.Split(srcaddr, "-")
	if dstaddrs[0] != srcaddrs[0] {
		RouteNetPath(srcaddr, dstaddr, leve)
	} else {
		RoutePath(srcaddr, dstaddr, leve)
	}
}
