package g

import (
	"bytes"
	"crypto/md5"
	//"encoding/hex"
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//"net/url"
	"encoding/json"

	//"strings"
	"time"
)

var (
	RootName string
	RootPass string
	Token    string
	SrcIP    []string
)

type Host struct {
	Id       string `json:"id"`
	HostName string `json:"hostname"`
}
type Hosts struct {
	HostArr []Host ` json:"host"`
}

type NetLink struct {
	SrcAddr string `json:"srcAddr"`
	SrcIPv4 string `json:"srcIPv4"`

	DstAddr string `json:"dstAddr"`
	DstIPv4 string `json:"dstIPv4"`

	FlowLevel int `json:flowLevel`
}

type Netinfo struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Ip   string `json:"ip"`
}

func readCfg() {
	RootName = Cfg.String("root_name")
	RootPass = Cfg.String("root_pass")
}
func timer() {
	timer1 := time.NewTicker(1800 * time.Second)
	for {
		select {
		case <-timer1.C:
			ClosedC1()
		}
	}
}

func ClosedC1() {
	for v := range C1 {
		delete(C1, v)
	}

}
func readIPFile() (src []string, err error) {
	f, e := os.Open("host.txt")
	if e != nil {
		err = e
		return
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		src = append(src, line)
	}
	return
}

func ReadHostInfo() (*Hosts, error) {
	hosts := &Hosts{}

	bytes, err := ioutil.ReadFile("host.json")
	if err != nil {
		return hosts, err
	}

	err = json.Unmarshal(bytes, hosts)
	if err != nil {
		fmt.Println(err)
		return hosts, err
	}
	return hosts, nil
}

func sendToken() {
	src, e := readIPFile()
	if e != nil {
		fmt.Println(e)
		return
	}
	SrcIP = src

	h := md5.New()
	current := time.Now().String()
	currents := []byte(current)
	h.Write([]byte(currents)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	Token = string(cipherStr)

	for _, host := range SrcIP {
		//fmt.Println(host)
		client := &http.Client{}
		req, err := http.NewRequest("POST", host, bytes.NewReader(cipherStr))
		//req1, err := http.NewRequest("POST", "http://10.0.0.179:8080/probe/ping", bytes.NewReader(jsonStr))
		if err != nil {
			// handle error
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Cookie", "name=anny")

		resp, err := client.Do(req)
		//resp1, err := client.Do(req1)

		defer resp.Body.Close()
		//defer resp1.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		//body1, err := ioutil.ReadAll(resp1.Body)
		//if err != nil {
		// handle error
		//}

		fmt.Printf("%s", body)
	}
	//fmt.Println(string(body1))
}

func netMap() error {
	netinfo := []Netinfo{}
	NetLinkMap = make(map[string]Netinfo)

	bytes, err := ioutil.ReadFile("netlink.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &netinfo)
	if err != nil {
		fmt.Println("hahha", err)
		return err
	}
	for _, v := range netinfo {
		NetLinkMap[v.Name] = v
	}
	//fmt.Println(NetLinkMap["1-1"])
	return nil
}
func NodeLinkSend(nodelink []byte, ip string) error {
	fmt.Println(ip)
	src := "http://127.0.0.1:9999/donetlink"
	client := &http.Client{}

	request, err := http.NewRequest("POST", src, bytes.NewReader(nodelink))
	if err != nil {
		fmt.Println("NewRequset error")
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	bodys := string(body)
	if bodys != "OK" {
		fmt.Println("send err")
		return err
	}

	return nil
}
