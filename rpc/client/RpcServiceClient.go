package client

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"os"
	"strconv"
	"time"
	"apromise-go-api/rpc"
)

var (
	HOST = "localhost"
	PORT = "9090"
)

func ClientStart() {

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	protocolFactory := thrift.NewTJSONProtocolFactory()

	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))

	if err != nil {
		fmt.Println(os.Stderr, "error resolving address", err)
	}

	useTransport, err := transportFactory.GetTransport(transport)

	client := rpc.NewRpcServiceClientFactory(useTransport, protocolFactory)

	if err := transport.Open(); err != nil {
		fmt.Println(os.Stderr, "error open socket to " + HOST + ":" + PORT, " ", err)
		os.Exit(1)
	}

	defer transport.Close()
	for i := 0; i < 10; i++ {
		paramMap := make(map[string]string)
		paramMap["a"] = "nihao yuhou"
		paramMap["b"] = "test " + strconv.Itoa(i + 1)
		r1, err := client.FunCall(time.Now().Unix(), "go client", paramMap)
		if (err != nil) {
			fmt.Println("调用错误",err)
		}
		fmt.Println(r1)
	}
}
