package rainrpc

import (
	"errors"
	"net"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"os"
	"strconv"
)

type Connection struct {
	Name   string               //链接名称
	Group  string               //客户端组
	Weight int                  //权重
	Client *InvokeServiceClient //服务客户端
}

type ConnectionFactory struct {
	Server      string                 //服务
	ServerList  []*RpcSerives          //服务列表
	ServerSize  int                    //服务个数
	Connections map[string]*Connection //服务链接信息
}

//初始化 链接 信息
var ConnectionFactorys = make(map[string]*ConnectionFactory);

//初始化服务链接工厂
func NewConnectionFactory() error {

	//如果依赖服务列表为空那么
	if len(Rpc_App_List) < 1 {
		return errors.New("no available server list!");
	}
	//如果链接池不为空那么直接返回链接池
	if len(ConnectionFactorys) > 1 {
		return nil
	}

	for _, data := range Rpc_App_List {
		factory := ConnectionFactory{}
		factory.Server = data.AppName
		factory.ServerList = data.Services
		factory.ServerSize = len(data.Services)
		ConnectionFactorys[factory.Server] = &factory
	}
	return nil
}

// 获取链接
func GetConnection(appName string) (*Connection, error) {

	if len(ConnectionFactorys) < 1 {
		return nil, errors.New("no available Connection!")
	}

	factory := ConnectionFactorys[appName]
	if factory == nil || factory.ServerSize < 1 {
		return nil, errors.New("no available " + appName + " Server!")
	}

	conMap := factory.Connections
	if len(conMap) < 1 {

		//这个地方路由占时放一下//默认去第一个进行链接
		rpcSerives := factory.ServerList[0]
		client := createConnection(rpcSerives.Ip, strconv.Itoa(rpcSerives.Port))

		connection := Connection{}
		connection.Name = rpcSerives.Ip
		connection.Client = client
		connection.Group = rpcSerives.App
		connection.Weight = 1//权重值目前都默认设置成1

		conn := make(map[string]*Connection)
		conn[rpcSerives.Ip] = &connection
		factory.Connections = conn
		return &connection, nil
	}

	for _, data := range conMap {
		//获取用户配置的路由方式进行路由选择链接
		//如果用户没有配置 默认随机链接
		return data, nil;
	}

	return nil, errors.New("no available connection!")

}

//创建链接
func createConnection(hosts string, port string) *InvokeServiceClient {

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(hosts, port))

	if err != nil {
		fmt.Println(os.Stderr, "error resolving address", err)
	}

	useTransport, _ := transportFactory.GetTransport(transport)

	client := NewInvokeServiceClientFactory(useTransport, protocolFactory)

	if errs := transport.Open(); errs != nil {
		fmt.Println(os.Stderr, "error open socket to", "172.18.53.24", 2181, errs)
		os.Exit(1)
	}

	return client;
	/*defer transport.Close()

	s, _ := client.Invoke(1, "2", "niceService", "sayNice", "")
	fmt.Println(s)

	s, _ = client.Invoke(1, "2", "niceService", "sayNice", "")
	fmt.Println(s)*/

}