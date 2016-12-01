package main

import (
	"bytes"
	"encoding/json"
	//	"flag"
	//"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	//	"os"
	//	"runtime"
)

func GetCurrentIp() {
	CurrentIp.PrivateIp = nil
	CurrentIp.PublicIp = nil
	CurrenHostIP = GetCompIp() //获取本机所有非自环ip
	if CurrenHostIP == nil {
		return
	}
	if Conf.HostType == "al0" {
		CurrentIp.PrivateIp = CurrenHostIP
		ipmsg := IpMsg{}
		ipmsg.Ip = Conf.Ip
		ipmsg.Isp = "阿里"
		CurrentIp.PublicIp = append(CurrentIp.PublicIp, ipmsg)
	} else {
		for _, v := range CurrenHostIP {
			if strings.HasPrefix(v, "10.") || strings.HasPrefix(v, "192.168") || strings.HasPrefix(v, "172.") {
				CurrentIp.PrivateIp = append(CurrentIp.PrivateIp, v)
			} else {

				isp, err := GetIsp(v)
				if err != nil {
					log.Printf("Clint abort! Cause:%v \n", err)
				}

				ipmsg := IpMsg{}
				ipmsg.Ip = v
				ipmsg.Isp = isp
				CurrentIp.PublicIp = append(CurrentIp.PublicIp, ipmsg)
			}
		}
	}

	err := SendNodeinfo(Srcpath + "/nodeinit")
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
		return
	}

}
func checkIP() {
	hostip := GetCompIp()
	var p string
	take := []string{}
	ok := false
	for _, v := range hostip {
		for _, p := range CurrenHostIP {
			if p == v {
				ok = true
				break
			}
		}
		if !ok {
			take = append(take, p)
			break
		}
	}
	if len(take) != 0 {
		GetCurrentIp()

	}
}
func GetCompIp() (src []string) {
	chack := make(map[string]int)
	addrs, errs := net.InterfaceAddrs()
	if errs != nil {
		log.Printf("Clint abort! Cause:%v \n", errs)
		return nil
	}

	for i, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				chack[ipnet.IP.String()] = i
			}
		}
	}
	for s, _ := range chack {
		src = append(src, s)
	}
	return
}

//func SendIP() {
//	sendips := SendIps{}
//	sendips.Addr = Conf.Addr
//	sendips.Ip = CurrentIp

//	path := "http://206.161.224.155:8787/ip"
//	bytereq, _ := json.MarshalIndent(&sendips, "", "")
//	body, err := SendHttpPost(bytereq, path)
//	if err != nil {
//		log.Printf("Client abort! Cause:%v \n", err)
//		return
//	}
//	if body == "" {
//		log.Printf("Client abort! Cause:send ip err")
//		return
//	}
//}
func GetIsp(ip string) (string, error) {

	//	var isp string
	//	isp = ""
	//	ispget := IspGet{}

	//	type Ip struct {
	//		ip string `json:"ip"`
	//	}
	//	ipsend := Ip{}
	//	ipsend.ip = ip

	//	byteip, err := json.MarshalIndent(&ipsend, "", "")
	//	if err != nil {
	//		log.Printf("Clint abort! Cause:%v \n", err)
	//		return isp, err
	//	}

	//	client := &http.Client{}
	//	path := "http://c.lonlife.org/api/ip/get"

	//	request, err := http.NewRequest("POST", path, bytes.NewReader(byteip))
	//	if err != nil {
	//		log.Printf("Client abort! Cause:%s \n", err)
	//		return isp, err
	//	}
	//	/request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//	//request.Header.Set("Cookie", "name=anny")

	//	resp, err := client.Do(request)
	//	if err != nil {
	//		log.Printf("Clint abort! Cause:%v \n", err)
	//		return isp, err
	//	}

	//	body, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		log.Printf("Clint abort! Cause:%v \n", err)
	//		return isp, err
	//	}
	//	fmt.Println("body", string(body))

	//	//err = json.Unmarshal(body, &resplnn)
	//	//fmt.Println(resplnn)

	//	err = json.Unmarshal(body, &ispget)
	//	if err != nil {
	//		log.Printf("Clint abort! Cause:%v \n", err, ip, string(body))
	//		return isp, err
	//	}
	//	if ispget.Msg != "success" {
	//		return isp, nil
	//	}
	//	isp = ispget.Result.Isp
	//	return isp, nil
	var isp string
	isp = ""
	ispget := IspGet{}

	type Ip struct {
		Ipp string `json:"ip"`
	}
	ipsend := Ip{}
	ipsend.Ipp = ip

	byteip, err := json.Marshal(&ipsend)
	if err != nil {
		log.Printf("Clint abort! Cause:%v \n", err)
		return isp, err
	}

	bodyip := bytes.NewBuffer([]byte(byteip))

	request, err := http.Post("http://c.lonlife.org/api/ip/get", "application/json;charset=utf-8", bodyip)
	if err != nil {
		log.Printf("Client abort! Cause:%s \n", err)
		return isp, err
	}
	result, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Clint abort! Cause:%v \n", err)
		return isp, err
	}

	err = json.Unmarshal(result, &ispget)
	if err != nil {
		log.Printf("Clint abort! Cause:%v \n", err, ip, string(result))
		return isp, err
	}
	if ispget.Msg != "success" {
		return isp, nil
	}
	isp = ispget.Result.Isp
	return isp, nil
}
