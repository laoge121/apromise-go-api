package rainrpc

import (
	"log"
	"github.com/coreos/etcd/client"
	"github.com/astaxie/beego"
)

var Cli client.KeysAPI

func init() {
	uris := beego.AppConfig.Strings("etcd.uris")
	cfg := client.Config{Endpoints:uris, Transport:client.DefaultTransport}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal("创建etcd客户端异常", err)
		panic("etcd init error! server uri is")
	}
	Cli = client.NewKeysAPI(c)
	//return kApi
}