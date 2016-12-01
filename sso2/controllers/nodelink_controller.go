package controllers

import (
	"encoding/json"
	"fmt"
	"sso2/g"
	"strconv"
)

func (this *RuleController) SetNodeLink() {

	this.Layout = "layout/admin.html"
	this.TplName = "netlink/table.html"
}

func (this *RuleController) DoNodeLink() {
	netlinks := []g.NetLink{}
	ids := this.GetStrings("ids")
	srcAddr := this.GetStrings("srcAddr")
	srcIPv4 := this.GetStrings("srcIPv4")
	dstAddr := this.GetStrings("dstAddr")
	dstIPv4 := this.GetStrings("dstIPv4")

	flowLevels := this.GetStrings("flowLevel")

	for k, _ := range ids {
		fmt.Println(k)
		netlink := g.NetLink{}

		srclink := g.NetLinkMap[srcAddr[k]]
		netlink.SrcAddr = srclink.Addr
		netlink.SrcIPv4 = srcIPv4[k]

		dstlink := g.NetLinkMap[dstAddr[k]]
		netlink.DstAddr = dstlink.Addr
		netlink.DstIPv4 = dstIPv4[k]

		flowLevelInt, err := strconv.Atoi(flowLevels[k])
		if err != nil {
			return
		}
		netlink.FlowLevel = flowLevelInt
		netlinks = append(netlinks, netlink)
	}
	//b, _ := json.Marshal(netlinks)
	b, _ := json.MarshalIndent(netlinks, "", "")

	link := g.NetLinkMap[srcAddr[0]]
	err := g.NodeLinkSend(b, link.Ip)
	if err != nil {
		return
	}

	this.Layout = "layout/admin.html"
	this.TplName = "netlink/table.html"
}
