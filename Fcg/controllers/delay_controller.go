package controllers

import (
	"Fcg/g"
	"encoding/json"
	"fmt"
	//"os"
	//"strconv"
	//	"strings"
	"strings"
)

func (this *NodeController) Delay() {
	delay := g.SendDelay{}

	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &delay)
	if err != nil {
		fmt.Println(" json error", err)
	}

	s := strings.Split(delay.Addr, "-") //解析发来的字符
	NetId := s[0]
	NodeId := s[1]

	if delay.NodeDelays != nil {
		for _, v := range delay.NodeDelays { //设置nodedelay
			if _, ok := g.RouteMap[NetId][NodeId][v.DestAddr]; ok {
				//sdlay := strconv.Itoa(v.AvgDelay)
				delaystr := v.AvgDelay
				tmpdelay := g.RouteMap[NetId][NodeId][v.DestAddr]
				tmpdelay.Delays = delaystr
				g.RouteMap[NetId][NodeId][v.DestAddr] = tmpdelay
			} else {
				fmt.Println("not fond " + v.DestAddr + " node delay")
			}
		}
	}

	if delay.NetDelays != nil {
		for _, v := range delay.NetDelays { //设置netdelay
			if _, ok := g.RouteMapNet[NetId][NodeId][v.DestAddr]; ok {
				//s := strconv.Itoa(v.AvgDelay)
				delaystr := v.AvgDelay
				tmpdelay := g.RouteMapNet[NetId][NodeId][v.DestAddr]
				tmpdelay.Delays = delaystr
				g.RouteMapNet[NetId][NodeId][v.DestAddr] = tmpdelay
			} else {
				fmt.Println("not fond " + v.DestAddr + " net delay")
			}

		}
	}
	this.Ctx.WriteString("")
}
