package main

import (
	//"encoding/json"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"strings"
)

func StartHTTP(addr string) error {
	http.HandleFunc("/donodelink", Donode)
	//	http.HandleFunc("/donetlink", Donet)
	//	http.HandleFunc("/dotransmit", DoTransmit)
	//	http.HandleFunc("/getmsg", Getmsg)
	return http.ListenAndServe(addr, nil)
}

func Donode(w http.ResponseWriter, r *http.Request) {
	_, _ = ioutil.ReadAll(r.Body)

	//path := "/home/L-Base/RouteNode/node_link.json"
	//	path := "node_link.json"
	//	//fmt.Println("hello this is client")
	//	err := WriteFile(body, path)
	//	if err != nil {
	//		log.Printf("Client abort! Cause:%s \n", err)
	//		b := []byte("NO")
	//		w.Write(b)
	//		return
	//	}
	//	b := []byte("OK")
	//	w.Write(b)
	return

}

func Donet(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	//path := "/home/L-Base/RouteNode/net_link.json"
	path := "net_link.json"
	log.Printf("Client abort! Cause:%s \n", body)
	err := WriteFile(body, path)
	if err != nil {
		fmt.Println(err)
		b := []byte("NO")
		w.Write(b)
		return
	}
	b := []byte("OK")
	w.Write(b)

}

//func DoTransmit(w http.ResponseWriter, r *http.Request) {
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Println(err)
//		b := []byte("NO")
//		w.Write(b)
//	}
//	path := "/home/L-Base/TransmitNode/bu.json"
//	log.Printf("Client abort! Cause:%s \n", body)
//	err = WriteFile(body, path)
//	if err != nil {
//		log.Printf("Client abort! Cause:%v \n", err)
//		b := []byte("NO")
//		w.Write(b)
//		return
//	}
//	b := []byte("OK")
//	w.Write(b)
//}
//func Getmsg(w http.ResponseWriter, r *http.Request) {
//	nodelink := []NodeLink{}
//	netlink := []NetLink{}
//	transmit := []TransmitNode{}
//	var netRoute string
//	var nodeRoute string
//	var delay string
//	req := Req{}

//	bytes, err := ioutil.ReadFile("/home/L-Base/RouteNode/node_link.json")
//	if err != nil {
//		log.Printf("Client abort! Cause:%v \n", err)
//	}
//	if err := json.Unmarshal(bytes, &nodelink); err != nil {
//		log.Printf("Client abort! Cause:%v \n", err)
//	}

//	bytes, err = ioutil.ReadFile("/home/L-Base/RouteNode/net_link.json")
//	if err != nil {
//		log.Printf("Client abort! Cause:%v \n", err)
//	}

//	if err := json.Unmarshal(bytes, &netlink); err != nil {
//		log.Printf("Client abort! Cause:%v \n", err)
//	}

//	bytes, err = ioutil.ReadFile("/home/L-Base/TransmitNode/bu.json")
//	if err != nil {
//		log.Printf("Client abort! Cause:%v \n", err)
//	}
//	if err := json.Unmarshal(bytes, &transmit); err != nil {
//		log.Printf("Client abort! Cause:%v \n", err)
//	}

//	bytes, err = ioutil.ReadFile("/dev/shm/lb_net_route.txt")
//	if err != nil {
//		log.Printf("Client abort! Cause:%v \n", err)
//	}
//	netRoute = string(bytes)

//	bytes, err = ioutil.ReadFile("/dev/shm/lb_node_route.txt")
//	if err != nil {
//		log.Printf("Client abort! Cause:%v \n", err)
//	}
//	nodeRoute = string(bytes)

//	bytes, err = ioutil.ReadFile("/dev/shm/lb_delay.txt")
//	if err != nil {
//		log.Printf("Client abort! Cause:%v \n", err)
//	}
//	delay = string(bytes)

//	req.Net = netlink
//	req.Node = nodelink
//	req.Transmit = transmit
//	req.NetRoute = netRoute
//	req.NodeRoute = nodeRoute
//	req.Delay = delay

//	bytereq, err := json.MarshalIndent(&req, "", "")
//	if err != nil {
//		b := []byte("err json req err")
//		w.Write(b)
//	}
//	w.Write(bytereq)

//}
