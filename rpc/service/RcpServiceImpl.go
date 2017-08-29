package service

import (
	"fmt"
	"time"
	"git.apache.org/thrift.git/lib/go/thrift"
	"apromise-go-api/rpc"
)

var (
	NetWorkSport = "127.0.0.1:9090"
)

type rpcService struct {

}

func (this *rpcService) FunCall(callTime int64, funCode string, paramMap map[string]string) (ret []string, err error) {
	fmt.Println("-->from client call:", time.Unix(callTime, 0).Format("2006-01-02 15:01:04"), funCode, paramMap)
	ret = append(ret, "key:" + paramMap["a"] + " value:" + paramMap["b"])
	return
}

func Start() {
	transportFactorys := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTJSONProtocolFactory()
	serverTransport, err := thrift.NewTServerSocket(NetWorkSport);
	if err != nil {
		fmt.Println("Error!", err)
	}
	handler := &rpcService{}
	process := rpc.NewRpcServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(process, serverTransport, transportFactorys, protocolFactory)
	fmt.Println("thrift server in", NetWorkSport)
	server.Serve()
}