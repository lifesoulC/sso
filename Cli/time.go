package main

import (
	"bytes"

	//"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
	"os/exec"
	//"runtime"
	"time"
)

func timeRoute() {
	timer1 := time.NewTicker(15 * time.Second)
	for {
		select {
		case <-timer1.C:
			checkCgf()      //检查路径信息 上报
			SendDelayinfo() //上报delay信息
		}
	}
}
func timeDelay() {
	timer1 := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timer1.C:
			checkdelay() //每一秒检查一次delay
		}
	}
}

func timepoll() {
	timer1 := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-timer1.C:
			Poll()
			//每一秒检查一次delay
		}
	}
}

func timeIp() {
	timer1 := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-timer1.C:
			checkIP() //每三分钟检查ＩＰ是否变化
		}
	}
}

func getTimer() {
	timer1 := time.NewTicker(30 * time.Minute)
	for {
		select {
		case <-timer1.C:
			getTime() //每三十分钟校对时钟
		}
	}
}

func httpGetTime() (string, error) {

	path := "http://nm.lbase.inc:9009/time"
	resp, err := http.Get(path)
	//request.Header.Set("Cookie", "test=John")
	if err != nil {
		fmt.Println(err)
	}
	//defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return string(body), err
	}

	return string(body), nil

}

func getTime() {
	var times string
	times = ""
	var err error
	begin_time := time.Now()
	//fmt.Println(begin_time)
	times, err = httpGetTime()
	if err != nil {
		log.Printf("Client abort! Cause:%v \n", err)
		//log.Printf("Client abort! Cause:%v \n", err)
	}
	end_time := time.Now()
	//fmt.Println(end_time)
	//fmt.Println("server time:", times)
	var dur_time time.Duration = end_time.Sub(begin_time)
	var elapsed_min float64 = dur_time.Minutes()
	intt := int(elapsed_min * 60)
	if intt >= 1 {
		return
	}

	if times != "" {
		//fmt.Println(times)
		in := bytes.NewBuffer(nil)
		cmd := exec.Command("sh")
		cmd.Stdin = in
		go func() {
			in.WriteString("date -s" + times)
		}()
		if err = cmd.Run(); err != nil {
			log.Printf("Client abort! Cause:%v \n", err)
			return
		}
	}
	return
}
