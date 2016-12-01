package main

import (
	//"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	//"flag"
	//"sync"
	"fmt"
	//"io/ioutil"
	"log"
	//"net/http"
	//"os"
	//"os/exec"
	//"runtime"
	"strconv"
	//"time"
	"strings"
)

func checkCgf() {

	pathnode := "/dev/shm/route/lb_node_route.txt"
	pathnet := "/dev/shm/route/lb_net_route.txt"

	bytenode, err := ReadFile(pathnode)
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
		return
	}
	strnode := string(bytenode)
	if strnode == "" {
		return
	}

	bytenet, err := ReadFile(pathnet)
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
		return
	}
	strnet := string(bytenet)
	if strnet == "" {
		log.Printf("Client abort! Cause:%v \n", err)

	}

	Md5NodeTmp := Md5(bytenode)
	Md5NetTmp := Md5(bytenet)

	if Md5Net == "" {
		Md5Net = Md5NetTmp
		SendRouteinfo(bytenet, "net")
	}
	if Md5Node == "" {
		Md5Node = Md5NodeTmp
		SendRouteinfo(bytenode, "node")
	}

	if Md5Net != Md5NetTmp {
		Md5Net = Md5NetTmp
		SendRouteinfo(bytenet, "net")
	}
	if Md5Node != Md5NodeTmp {
		Md5Node = Md5NodeTmp
		SendRouteinfo(bytenode, "node")
	}

}

func checkdelay() {
	path_node_delay := "/dev/shm/lb_node_delay.txt"
	path_net_delay := "/dev/shm/lb_net_delay.txt"

	//path_node_delay := "lb_node_delay.txt"
	//path_net_delay := "lb_net_delay.txt"

	//计算nodedelay
	bytenode_delay, err := ReadFile(path_node_delay)
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
	}
	strnode_delay := string(bytenode_delay)
	//fmt.Println("delay_node:", strnode_delay)

	Lstrs := strings.SplitN(strnode_delay, "\n", -1) //得出每行的信息 1-1 118.187.3.6 -- 1-2 123.57.6.17: 3
	for _, v := range Lstrs {
		if v == "" {
			break
		}
		Wstrs := strings.SplitN(v, " ", -1) //得出每个单词的信息 "1-1"  "118.187.3.6"

		if Wstrs[0] == Conf.Addr {
			Mutex.Lock()
			Nodedelaymap[Wstrs[3]] = append(Nodedelaymap[Wstrs[3]], Wstrs[5])
			Mutex.Unlock()
		}
	}
	//计算netdelay
	bytenet_delay, err := ReadFile(path_net_delay)
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
		return
	}
	strnet_delay := string(bytenet_delay)
	if strnet_delay == "" {
		typestr := []string{}
		typestr = append(typestr, "0")
		Mutex.Lock()
		Netdelaymap["type"] = typestr
		Mutex.Unlock()
		return
	}
	Mutex.Lock()
	Netdelaymap["type"] = append(Netdelaymap["type"], "1")
	Mutex.Unlock()
	NetLstrs := strings.SplitN(strnet_delay, "\n", -1) //得出每行的信息 1-1 118.187.3.6 -- 1-2 123.57.6.17: 3
	for _, v := range NetLstrs {
		if v == "" {
			break
		}
		Wstrs := strings.SplitN(v, " ", -1) //得出每个单词的信息 "1-1"  "118.187.3.6"

		if Wstrs[0] == Conf.Addr {
			Mutex.Lock()
			Netdelaymap[Wstrs[3]] = append(Netdelaymap[Wstrs[3]], Wstrs[5])
			Mutex.Unlock()
		}
	}
	return

}
func SendDelayinfo() {
	fmt.Println("sendDelayinfo cfg 132")
	path := "http://nm.lbase.inc:9009/delay"
	Senddelay.Addr = Conf.Addr
	Senddelay.Name = Conf.Name
	for k, v := range Nodedelaymap {
		tmpdelay := Delay{}
		var sum int = 0
		var cont int = 0
		for _, sdelay := range v {
			i, err := strconv.Atoi(sdelay)
			if err != nil {
				log.Printf("Client abort! Cause:%v \n", err)
				continue
			}
			sum = sum + i
			cont++
		}
		tmpdelay.DestAddr = k
		tmpdelay.AvgDelay = sum / cont

		Senddelay.NodeDelays = append(Senddelay.NodeDelays, tmpdelay)

	}
	Mutex.Lock()
	if v, ok := Netdelaymap["type"]; ok {
		fmt.Println(v)
		if v[0] == "0" {
			Senddelay.NetDelays = nil
		} else {
			for k, v := range Netdelaymap {
				if k == "type" {
					continue
				}
				tmpdelay := Delay{}
				var sum int = 0
				var cont int = 0
				for _, sdelay := range v {
					i, err := strconv.Atoi(sdelay)
					if err != nil {
						log.Printf("Client abort! Cause:%v \n", err)
						continue
					}
					sum = sum + i
					cont++
				}
				tmpdelay.DestAddr = k
				tmpdelay.AvgDelay = sum / cont

				Senddelay.NetDelays = append(Senddelay.NetDelays, tmpdelay)

			}
		}
	}
	Mutex.Unlock()

	bytSenddelay, err := json.MarshalIndent(&Senddelay, "", "")
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
		return
	}

	fmt.Println(Senddelay.NodeDelays)
	fmt.Println(Senddelay.NetDelays)

	deletemap() //清空map

	Senddelay.NetDelays = nil
	Senddelay.NodeDelays = nil

	_, err = SendHttpPost(bytSenddelay, path) //发送出去
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
		return
	}
}
func Md5(file []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(file)
	cipherStr := md5Ctx.Sum(nil)

	Mdfile := hex.EncodeToString(cipherStr)
	return Mdfile
}

func SendRouteinfo(filebyte []byte, Type string) {
	path := "http://nm.lbase.inc:9009/lbroute"
	sendaddr := SendRoute{}
	sendaddr.Addr = Conf.Addr
	sendaddr.Content = filebyte
	sendaddr.Type = Type
	bytereq, _ := json.MarshalIndent(&sendaddr, "", "")
	_, err := SendHttpPost(bytereq, path)
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
		return
	}
	//log.Printf("Client abort! Cause:%v \n", string(body))
}
