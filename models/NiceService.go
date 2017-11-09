package models

import (
	"apromise-go-api/rainrpc"
	"fmt"
	"encoding/json"
)

type Person struct {
	Id   int `json:"id"`
	Name string `json:"name"`
	Ret  map[string]string `json:"ret"`
}

const appName = "raining"

const server = "niceService"

func SayNice() string {
	rainrpc.NewConnectionFactory()
	conn, err := rainrpc.GetConnection(appName)
	if err != nil {
		fmt.Println(err)
	}
	requst:=rainrpc.RainRequest{};
	ret, err := conn.Execute(&requst)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return ret.ResponseHead.Message
}

func SayNice1(p Person) string {
	conn, err := rainrpc.GetConnection(appName)
	if err != nil {
		fmt.Println(err)
	}
	_, e := json.Marshal(p)
	if e != nil {
		return ""
	}
	requst:=rainrpc.RainRequest{};
	ret, err := conn.Execute(&requst)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	return ret.ResponseHead.Message
}
