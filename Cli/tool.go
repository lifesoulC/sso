package main

import (
	"bytes"
	//	"crypto/md5"
	//	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//	"os/exec"
	"runtime"
	//	"time"
)

func SendNodeinfo(srcpath string) error {

	sendinfo := Sendinfo{}
	sendinfo.Confginfo = Conf
	sendinfo.Ips = CurrentIp

	bytSendinfo, err := json.MarshalIndent(&sendinfo, "", "")
	fmt.Println(string(bytSendinfo))
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
		return err
	}
	fmt.Println("sendinfo:", sendinfo)

	_, err = SendHttpPost(bytSendinfo, srcpath)
	if err == nil {
		fmt.Println("send ok")
	} else {
		return err
	}

	return nil
}

func ReadFile(path string) ([]byte, error) {
	M.RLock()
	bytes, err := ioutil.ReadFile(path)
	M.RUnlock()
	if err != nil {
		return nil, err
	}
	return bytes, nil

}

//func ChecFile() {
//	path := "/home/L-Base/RouteNode/node_link.json"
//	bytes, err := ReadFile(path)
//	if err != nil {
//		fmt.Println(err)
//	}
//	s := string(bytes)
//	if s == "" {
//		var path string
//		if T == 1 {
//			path := "http://183.60.189.27:9990/time"
//		}
//		if T == 2 {
//			path := "http://206.161.224.155:9990/time"
//		}
//		//srcpath := "http://206.161.224.155:9990/chicklips"
//		err = SendAddr(path)
//		if err != nil {
//			fmt.Println(err)
//			//return
//		}
//	}
//}
func loginit() error {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	//set logfile Stdout
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		return logErr
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return nil
}

func SendHttpPost(Content []byte, path string) ([]byte, error) { //发送成功 返回服务器返回信息 失败 返回空

	b := []byte{}
	client := &http.Client{}

	request, err := http.NewRequest("POST", path, bytes.NewReader(Content))
	if err != nil {
		log.Printf("Client abort! Cause:%s \n", err)
		return b, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Cookie", "test=John")

	resp, err := client.Do(request)
	if err != nil {
		return b, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return b, err
	}

	return body, nil
}

func GetHostType() error {

	path := "/etc/lbase.conf"
	//path := "lbase.txt"
	bytest, err := ReadFile(path)
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
		return err
	}
	if err := json.Unmarshal(bytest, &Conf); err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
		return err
	}

	return nil

}

func deletemap() {
	for k, _ := range Nodedelaymap {
		Mutex.Lock()
		delete(Nodedelaymap, k)
		Mutex.Unlock()
	}
	//fmt.Println("delay Node", Nodedelaymap)
	for k, _ := range Netdelaymap {
		Mutex.Lock()
		delete(Netdelaymap, k)
		Mutex.Unlock()
	}
	//	fmt.Println("delay Node", Netdelaymap)
}
func Poll() {
	addr := Conf.Addr
	sendpoll := Sendpoll{}
	sendpoll.Addr = addr

	bytepoll, err := json.MarshalIndent(&sendpoll, "", "")
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
		return
	}

	body, err := SendHttpPost(bytepoll, Srcpath+"/poll")
	if err == nil {
		fmt.Println("收到node")
		fmt.Println("tool 160 body :", string(body))
		reqpoll := Reqpoll{}
		if err := json.Unmarshal(body, &reqpoll); err != nil {
			fmt.Println(err)
			return
		} else {
			if len(reqpoll.ContentNode) != 0 {
				err := WriteFile(reqpoll.ContentNode, "node_link.json")
				if err != nil {
					log.Printf("Client abort! Cause:%v \n", err)
					return
				}
			} else {
				if len(reqpoll.ContentNet) != 0 {
					err := WriteFile(reqpoll.ContentNet, "net_link.json")
					if err != nil {
						log.Printf("Client abort! Cause:%v \n", err)
						return
					}
				}
			}
		}
	}
}
