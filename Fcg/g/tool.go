package g

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"
)

var (
	LogFileName = flag.String("log", "cserver.log", "Log file name")
)

func AddNodemapinfo(nodemapinfo Nodelinkmap) { //？？添加进入 Nodemapinfo
	node := nodemapinfo.Node
	net := nodemapinfo.Net
	if nodemap, ok := NodeMap[nodemapinfo.Net]; ok { //查看是否有该 node节点信息  map[net]map[node] Nodemap
		if nodeinfo, ok := nodemap[nodemapinfo.Node]; ok {

			for _, oldIpprivat := range nodeinfo.Ipprivate {
				for _, v := range nodemapinfo.Ipprivate {
					if v.Ip == oldIpprivat.Ip {
						nodemapinfo.Status = oldIpprivat.Status
					}
				}
			}

			for _, oldIppublic := range nodeinfo.Ippublic {
				for _, v := range nodemapinfo.Ippublic {
					if v.Ip == oldIppublic.Ip {
						nodemapinfo.Status = oldIppublic.Status
					}
				}
			}
			nodemapinfo.Status = nodeinfo.Status

			nodemap[node] = nodemapinfo
			NodeMap[net] = nodemap
		} else {
			nodemap[node] = nodemapinfo
			NodeMap[net] = nodemap
		}

	} else {
		nodemap := make(map[string]Nodelinkmap)
		nodemap[node] = nodemapinfo
		NodeMap[net] = nodemap

		//添加进 Nodemap
	}
	fmt.Println("tool 56:NodeMap", NodeMap)
}
func InitCfg(net string) {
	nodemap := RouteMap[net]
	fmt.Println("tool 60 RouteMap[net]:", nodemap)
	//	fmt.Println(nodemap)
	var b []byte
	for node, destmap := range nodemap {
		respcfgs := []RespCfg{}
		for _, nodemapinfo := range destmap {
			respcfg := RespCfg{}
			respcfg.SrcAddr = nodemapinfo.From.NodeId
			respcfg.SrcIPv4 = nodemapinfo.From.Ip

			respcfg.DstAddr = nodemapinfo.To.NodeId
			respcfg.DstIPv4 = nodemapinfo.To.Ip

			respcfg.FlowLevel = 1
			if respcfg.DstAddr == "" {
				continue
			}
			respcfgs = append(respcfgs, respcfg)
		}
		if len(respcfgs) == 0 {
			str := "[]"
			b = []byte(str)
			if _, ok := PollMapNode[net]; ok {

				PollMapNode[net][node] = b

			} else {
				dstmap := make(map[string][]byte)
				dstmap[node] = b
				PollMapNode[net] = dstmap
			}
		} else {
			//	fmt.Println(respcfgs)
			b, _ = json.MarshalIndent(respcfgs, "", "")
			if _, ok := PollMapNode[net]; ok {
				PollMapNode[net][node] = b
			} else {
				dstmap := make(map[string][]byte)
				dstmap[node] = b
				PollMapNode[net] = dstmap
			}
		}

		//		ip := NodeMap[net][node].Ip
		//		fmt.Println("ip 92:", net+"-"+node, ip)
		//		fmt.Println("tool 91 :下发的数据", string(b))

		//		path := "http://" + ip + ":9900/donodelink"
		//		reqbody, err := HttpPostCfg(b, path) //下发
		//		if (err != nil) || (reqbody == "") {
		//			fmt.Println("tool 94 sedn fcg err", err)
		//		}
		//		strbody := string(reqbody)
		//		if strbody == "OK" {
		//			fmt.Println(net + "-" + node + "success")
		//			return
		//		}
	}
	fmt.Println("PollMap :", PollMapNode)

}

func NilCfg(net string, node string) {
	//nodemap := RouteMap[net]
	//	fmt.Println(nodemap)
	//var b []byte
	strs := "[]"
	b := []byte(strs)
	PollMapNode[net][node] = b

	//	ip := NodeMap[net][node].Ip
	//	fmt.Println("下发", net+"-"+node)
	//	fmt.Println("ip", ip)
	//	fmt.Println(nodemap)
	//	fmt.Println(strs)

	//	path := "http://" + ip + ":9900/donodelink"
	//	reqbody, err := HttpPostCfg(b, path) //下发
	//	if (err != nil) || (reqbody == "") {
	//		fmt.Println("node_controller 121", err)
	//	}
	//	strbody := string(reqbody)
	//	if strbody == "OK" {
	//		fmt.Println(net + "-" + node + "success")
	//		return
	//	}
}

//func HttpPostCfg(Content []byte, path string) (string, error) { //发送成功 返回服务器返回信息 失败 返回空

//	s := ""
//	client := &http.Client{}
//	fmt.Println("too 143 content:", string(Content))
//	request, err := http.NewRequest("POST", path, bytes.NewReader(Content))
//	if err != nil {
//		log.Printf("Client abort! Cause:%s \n", err)
//		return s, err
//	}
//	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	request.Header.Set("Cookie", "name=anny")

//	resp, err := client.Do(request)
//	if err != nil {
//		return s, err
//	}

//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return s, err
//	}
//	s = string(body)
//	if s == "" {
//		return s, err
//	}
//	if err != nil {
//		return s, err
//	}
//	defer resp.Body.Close()
//	return s, nil

//}
func HttpPostCfg(Content []byte, path string) (string, error) {

	resp, err := HttpGetFromIP(path, Content)
	if err != nil {
		fmt.Println("170 tool :", err)
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("tool 168 postcfg:", err)
		return "", err
	}
	fmt.Println("resp:", string(body))
	return string(body), nil
}

func Loginit() error {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	//set logfile Stdout
	logFile, logErr := os.OpenFile(*LogFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		return logErr
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return nil
}
func HttpGetFromIP(url string, Content []byte) (*http.Response, error) {
	req, _ := http.NewRequest("POST", url, bytes.NewReader(Content))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				//本地地址  ipaddr是本地外网IP
				lAddr, err := net.ResolveTCPAddr(netw, "183.60.189.27"+":8989")
				if err != nil {
					return nil, err
				}
				//被请求的地址
				rAddr, err := net.ResolveTCPAddr(netw, addr)
				if err != nil {
					return nil, err
				}
				conn, err := net.DialTCP(netw, lAddr, rAddr)
				if err != nil {
					return nil, err
				}
				deadline := time.Now().Add(2 * time.Second)
				conn.SetDeadline(deadline)
				return conn, nil
			},
		},
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36")
	return client.Do(req)
}
