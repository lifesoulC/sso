package g

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

type IPinfo struct {
	Ip     string `json:"ip"`
	Status int    `json:"status"`
	Isp    string `json:"isp"`
}

//发送给前段的信息 node
type Response struct {
	Status int    `json:"status"`
	Datas  []Data `json:"data"`
}

//节点的具体信息
type Data struct {
	Flow     string   `json:"flow"`
	Id       string   `json:"id"`
	Ipinfo   []IPinfo `json:"ip"`
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Net      string   `json:"net"`
	Status   int      `json:"status"`
}

type Nodelinkmap struct {
	Flow      string   `json:"flow"`      //流量
	Id        string   `json:"id"`        //“1-3”
	Ippublic  []IPinfo `json:"ippublic"`  // “公有IP”
	Ipprivate []IPinfo `json:"ipprivate"` //“私有IP”
	Name      string   `json:"name"`      //"bjvm"
	Location  string   `json:"location"`  //"北京"
	HostType  string   `json:"hosttype"`  //“vm”
	Net       string   `json:"net"`       //"1"
	Node      string   `json:"node"`
	Status    int      `json:"status"` // 1 启动  0 关闭
	Ip        string   `json:"ip"`
}
type Nodesmpinfo struct {
	NodeId   string `json:"nodeId"` //"1-1"
	Location string `json:"location"`
	Ip       string `json:"ip"`
	Isp      string `json:"isp"` //"北京"
}
type RoutlinkMap struct {
	From   Nodesmpinfo `json:"from"`   //源地址
	To     Nodesmpinfo `json:"to"`     //目的地址
	Level  int         `json:"level"`  //浏量等级
	Delays int         `json:"delay"`  //延迟  "20ms"
	Lost   string      `json:"lost"`   //丢包率
	Status int         `json:"status"` //默认为0表示自动生成的  1 为手动修改
}

type RoutelinkMapNet struct {
	From   Nodesmpinfo `json:"from"`   //源地址
	To     Nodesmpinfo `json:"to"`     //目的地址
	Level  int         `json:"level"`  //浏量等级
	Delays int         `json:"delay"`  //延迟  "20ms"
	Lost   string      `json:"lost"`   //丢包率
	Status int         `json:"status"` //默认为0表示自动生成的  1 为手动修改
}

type RespCfg struct {
	SrcAddr string `json:"srcAddr"`
	SrcIPv4 string `json:"srcIPv4"`

	DstAddr string `json:"dstAddr"`
	DstIPv4 string `json:"dstIPv4"`

	FlowLevel int `json:"flowLevel"`
}
type LbRoute struct {
	SrcAddr string
	SrcIp   string

	DstAddr string
	Dstip   string
}
