package main

//type NodeLink struct {
//	SrcAddr string `json:"srcAddr"`
//	SrcIPv4 string `json:"srcIPv4"`

//	DstAddr string `json:"dstAddr"`
//	DstIPv4 string `json:"dstIPv4"`

//	FlowLevel int `json:"flowLevel"`
//}

//type NetLink struct {
//	SrcAddr string `json:"srcAddr"`
//	SrcIPv4 string `json:"srcIPv4"`

//	DstAddr string `json:"dstAddr"`
//	DstIPv4 string `json:"dstIPv4"`

//	FlowLevel int `json:flowLevel`
//}

//type TransmitNode struct {
//	BuId        int    `json:"buID"`
//	BuName      string `json:"buName"`
//	AdapterPort int    `json:"adapterPort"`
//}

//type Req struct {
//	Node      []NodeLink     `json:"nodeLink"`
//	Net       []NetLink      `json:"netLink"`
//	Transmit  []TransmitNode `json:"transmitNode"`
//	NodeRoute string         `json:"nodeRoute"`
//	NetRoute  string         `json:"netRoute"`
//	Delay     string         `json:"delay"`
//}

type Confg struct {
	Addr     string `json:"addr"`
	Name     string `json:"name"`
	Location string `json:"location"`
	HostType string `json:"hosttype"`
	Ip       string `json:"ip"`
}
type Delay struct {
	DestAddr string `json:"destaddr"`
	AvgDelay int    `json:"avgdelay"`
}

type SendDelay struct {
	Addr       string  `json:"addr"`
	Name       string  `json:"name"`
	NodeDelays []Delay `json:"delays"`
	NetDelays  []Delay `json:"netdelays"`
}

type SendRoute struct {
	Addr    string `json:"addr"`
	Type    string `json:"type"`
	Content []byte `json:"content"`
}
type IPs struct {
	PublicIp  []IpMsg  `json:"publicip"`
	PrivateIp []string `json:"privateip"`
}
type ISP struct {
	Isp string `json:"isp"`
}

type IspGet struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result ISP    `json:"result"`
}

type IpMsg struct {
	Ip  string `json:"ip"`
	Isp string `json:"isp"`
}
type SendIps struct {
	Addr string `json:"addr"`
	Ip   IPs    `json:"Ip"`
}
type Sendinfo struct {
	Confginfo Confg `json:"confginfo"`
	Ips       IPs   `json:"ips"`
}
type Sendpoll struct {
	Addr string `json:"addr"`
}
type Reqpoll struct {
	//Types   string `json:"type"`
	ContentNet  []byte `json:"contentnet"`
	ContentNode []byte `json:"contentnode"`
}
