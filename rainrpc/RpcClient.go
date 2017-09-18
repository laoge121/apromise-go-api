package rainrpc

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"os"
)

//rpc 客户端服务
type rpcClient struct {
}

func ClientInvoke() {

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("localhost", "2181"))

	if err != nil {
		fmt.Println(os.Stderr, "error resolving address", err)
	}

	useTransport, _ := transportFactory.GetTransport(transport)

	client := NewInvokeServiceClientFactory(useTransport, protocolFactory)

	if errs := transport.Open(); errs != nil {
		fmt.Println(os.Stderr, "error open socket to", "172.18.53.24", 2181, errs)
		os.Exit(1)
	}
	defer transport.Close()

	s, _ := client.Invoke(1, "2", "niceService", "sayNice", "")
	fmt.Println(s)

	s, _ = client.Invoke(1, "2", "niceService", "sayNice", "")
	fmt.Println(s)
}

func test() {
}