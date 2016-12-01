package main

import (
	"fmt"
	//"os"
	"flag"
	"log"
	"sync"
)

var (
	logFileName = flag.String("log", "cClient.log", "Log file name")
)
var M *sync.RWMutex
var Md5Node string
var Md5Net string
var Mutex *sync.Mutex

//var T int
var Srcpath string
var Conf Confg

var Netdelaymap map[string][]string
var Nodedelaymap map[string][]string
var Senddelay SendDelay
var CurrentIp IPs
var CurrenHostIP []string

func main() {

	Md5Node = ""
	Md5Net = ""
	Conf = Confg{}
	Senddelay = SendDelay{}
	Netdelaymap = make(map[string][]string)
	Nodedelaymap = make(map[string][]string)
	M = new(sync.RWMutex)
	Mutex = new(sync.Mutex)
	CurrentIp = IPs{}

	var err error

	err = loginit()
	if err != nil {
		fmt.Println("loginit err")
	}

	port := 9900
	listenPort := fmt.Sprintf(":%d", port)
	fmt.Println("listen port", port)

	Srcpath = "http://nm.lbase.inc:9009"
	err = GetHostType() //获取本机的信息
	if err != nil {
		return
	}
	go timeRoute() //检测配置文件是否有变更  1：上报delay
	go timeDelay() //更新delay
	go timeIp()    //更新ip

	//getTime() //更新时间 每30min 更新一次  时间设置

	//go getTimer() //30分钟更新时间一次

	GetCurrentIp() //启动时出发一次
	go timepoll()

	err = StartHTTP(listenPort)
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
	}

}
