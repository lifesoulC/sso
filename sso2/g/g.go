package g

import (
	"log"

	//"fmt"

	"time"

	"github.com/astaxie/beego"
	mgo "gopkg.in/mgo.v2"
)

var Cfg = beego.AppConfig

var (
	PersonDB *mgo.Database
	Person   *mgo.Collection
	Role     *mgo.Collection
)
var HostArry *Hosts
var Ops int64 = 0
var M1 map[string]string
var T1 map[string]time.Time //存储超时信息
var C1 map[string]int
var NetLinkMap map[string]Netinfo

func InitEnv() {
	db_type := Cfg.String("db_type")
	db_host := Cfg.String("db_host")
	db_port := Cfg.String("db_port")
	db_name := Cfg.String("db_name")
	db_user := Cfg.String("db_user")
	db_pass := Cfg.String("db_pass")

	db_collection := Cfg.String("db_collection")
	db_role := Cfg.String("db_role")
	M1 = make(map[string]string)
	T1 = make(map[string]time.Time)
	C1 = make(map[string]int)
	if db_type == "mongo" {
		mongo_session, err := mgo.Dial("mongodb://" + db_user + ":" + db_pass + "@" + db_host + ":" + db_port) //链接主机
		//mongo_session, err := mgo.Dial(db_host + ":" + db_port) //链接主机
		if err != nil {
			log.Fatal(err)
		}

		PersonDB = mongo_session.DB(db_name) //得到表
		Person = PersonDB.C(db_collection)   //链接rules表
		Role = PersonDB.C(db_role)           //链接 role表

		mongo_session.SetMode(mgo.Monotonic, true)
	} else {
		log.Fatal("DB type is not mongo")
	}
	var err error
	err = netMap()
	if err != nil {
		log.Fatal(err)
	}
	readCfg()

	HostArry, err = ReadHostInfo()
	if err != nil {
		log.Fatal(err)
	}
	go timer()
}
