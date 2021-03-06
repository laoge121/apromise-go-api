package rainrpc

import (
	"github.com/astaxie/beego"
	"log"
	"context"
	"strings"
	"strconv"
)

//rpc 服务方法
type RpcMethod struct {
	MethodName  string
	Parametes   string
	MethodParam map[string]string
}

//rpc 服务列表
type RpcServer struct {
	ClassName   string
	ClassMethod []*RpcMethod
}

//rpc 服务地址
type RpcSerives struct {
	App     string
	Ip      string
	Port    int
	Servers []*RpcServer
}

type RpcApp struct {
	AppName  string        //应用名称
	Services []*RpcSerives //服务列表
}

//订阅的 服务列表
var Rpc_App_List = make([]*RpcApp, 1)

func LoadService() {
	servives := beego.AppConfig.Strings("rpc.services")
	length := 0
	for _, service := range servives {
		rpcApp := RpcApp{}
		rpcApp.AppName = service

		path := "/rain-server-raining" + "/" + service;
		resp, errs := Cli.Get(context.Background(), path, nil)
		if ( nil != errs) {
			log.Fatal("load server list error, no available server!")
		}
		nodeData := resp.Node.Value
		//如果数据不是正确ip：port 格式那么就继续吓一条数据处理
		if (!strings.Contains(nodeData, ":")) {
			log.Fatal("hosts info error!", nodeData)
			continue
		}
		hosts := strings.Split(nodeData, ";")
		var Rpc_server_list = make([]*RpcSerives, len(hosts))
		size := 0
		for _, str := range hosts {
			host := strings.Split(str, ":")
			rpcs := RpcSerives{}
			ip := host[0]
			rpcs.Ip = ip
			port, _ := strconv.Atoi(host[1])
			rpcs.Port = port;
			Rpc_server_list[size] = &rpcs
			size++
		}
		rpcApp.Services = Rpc_server_list
		Rpc_App_List[length] = &rpcApp
	}
}

//加载服务
/*
func LoadService() {

	//设置基础的服务列表
	servers := beego.AppConfig.Strings("rpc.services")
	appNum := 0;
	for _, serverPath := range servers {

		serverNames := strings.Split(serverPath, "\\.");
		if (len(serverNames) < 3) {
			continue
		}
		rpcApp := RpcApp{}
		rpcApp.AppName = serverNames[2]

		basePath := "/rain-server-raining" + "/" + serverPath;
		resp, errs := Cli.Get(context.Background(), basePath, nil)
		if errs != nil {
			log.Fatal("load server list error, no available server!")
		}
		sort.Sort(resp.Node.Nodes)

		//指定服务的单机app信息
		var Rpc_server_list = make([]*RpcSerives, 1)

		///rain-server-raining/172.18.53.24:2181
		snum := 0
		for _, n := range resp.Node.Nodes {
			serverHost := n.Key
			serverHost = strings.Replace(serverHost, basePath+"/", "", -1)
			hosts := strings.Split(serverHost, ":")
			rpcs := RpcSerives{}
			ip := hosts[0]
			rpcs.Ip = ip
			port, _ := strconv.Atoi(hosts[1])
			rpcs.Port = port;
			res, err := Cli.Get(context.Background(), n.Key, nil)
			if err != nil {
				log.Fatal("load server interface error, no available server interface!")
				continue
			}
			rpcss := make([]*RpcServer, res.Node.Nodes.Len())
			///rain-server-raining/172.18.53.24:2181/hellowService
			size := 0
			for _, nn := range res.Node.Nodes {
				sName := nn.Key
				sName = strings.Replace(sName, n.Key+"/", "", -1)
				rpcsr := RpcServer{}
				rpcsr.ClassName = sName

				re2, er2 := Cli.Get(context.Background(), nn.Key, nil)
				if er2 != nil {
					log.Fatal("load server interface method error, no available server interface method!")
					continue
				}
				rpcsm := make([]*RpcMethod, re2.Node.Nodes.Len())
				num := 0
				for _, nnn := range re2.Node.Nodes {
					mName := nnn.Key
					mName = strings.Replace(mName, nn.Key+"/", "", -1)
					names := strings.Split(mName, "@")
					rM := RpcMethod{}
					rM.MethodName = names[0]
					rM.Parametes = names[1]
					data := make(map[string]string)
					json.Unmarshal([]byte(nnn.Value), &data)
					rM.MethodParam = data
					rpcsm[num] = &rM
					num++
					fmt.Println("method info:", rM)
				}

				fmt.Println("service method list:", rpcsm)

				rpcsr.ClassMethod = rpcsm
				rpcss[size] = &rpcsr
				size++
			}
			rpcs.Servers = rpcss
			Rpc_server_list[snum] = &rpcs
		}
		rpcApp.Services = Rpc_server_list
		Rpc_App_List[appNum] = &rpcApp
		appNum++
	}
	for _, app := range Rpc_App_List {
		fmt.Println("服务应用信息:::", "---", app.AppName, "---")
		for _, data := range app.Services {
			fmt.Println("服务地址:::", "---", data.Ip, "---", data.Port)
			for _, data1 := range data.Servers {
				fmt.Println("服务名称:::", "---", data1.ClassName, "---")
				for _, data2 := range data1.ClassMethod {
					fmt.Println("服务方法列表:::", "---", data2.MethodName, "---", data2.Parametes, "---", data2.MethodParam)
				}
			}
		}
	}
}
*/
