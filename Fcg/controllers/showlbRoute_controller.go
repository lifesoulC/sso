package controllers

import (
	"Fcg/g"
	"encoding/json"
	"fmt"
	//"os"
	//	"strconv"
	"strings"
)

func (this *NodeController) ShowLbRoute() {
	type Response struct {
		Status int    `json:"status"`
		Data   string `json:"data"`
	}
	response := Response{}
	types := this.GetString("type")
	addr := this.GetString("addr")

	if types == "net" {
		if v, ok := g.LbRouteNet[addr]; ok {
			response.Status = 1
			response.Data = v
		} else {
			response.Status = 0
			response.Data = "can not find this addr route" + addr
		}
	} else {
		if v, ok := g.LbRouteNode[addr]; ok {
			response.Status = 1
			response.Data = v
		} else {
			response.Status = 0
			response.Data = "can not find this addr route" + addr
		}
	}

	this.Data["json"] = reqpoll
	this.ServeJSON()
}
