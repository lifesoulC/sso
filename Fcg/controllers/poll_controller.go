package controllers

import (
	"Fcg/g"
	"encoding/json"
	"fmt"
	//"os"
	//	"strconv"
	"strings"
)

func (this *NodeController) Poll() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	type Sendpoll struct {
		Addr string `json:"addr"`
	}
	type Reqpoll struct {
		//Types   string `json:"type"`
		ContentNode []byte `json:"contentnode"`
		ContentNet  []byte `json:"contentnet"`
	}
	sendpoll := Sendpoll{}
	reqpoll := Reqpoll{}

	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &sendpoll)
	if err != nil {
		fmt.Println(" json error", err)
	}

	addr := sendpoll.Addr

	s := strings.Split(addr, "-") //解析发来的字符
	NetId := s[0]
	NodeId := s[1]

	if v, ok := g.PollMapNode[NetId][NodeId]; ok {
		if len(v) != 0 {
			reqpoll.ContentNode = v
			fmt.Println("poll 39", reqpoll)
			g.PollMapNode[NetId][NodeId] = []byte("")
		} else {
			reqpoll.ContentNode = []byte("")
		}
	}

	if v, ok := g.PollMapNet[NetId][NodeId]; ok {
		if len(v) != 0 {
			reqpoll.ContentNet = v
			fmt.Println("poll 66", reqpoll)
			g.PollMapNet[NetId][NodeId] = []byte("")
		} else {
			reqpoll.ContentNet = []byte("")
		}
	}

	fmt.Println("content poll 72", string(reqpoll.ContentNet))
	fmt.Println("content poll 72", string(reqpoll.ContentNode))
	this.Data["json"] = reqpoll
	this.ServeJSON()

}
