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
	ret, err := conn.Execute(server, "sayNice", "")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return ret
}

func SayNice1(p Person) string {
	conn, err := rainrpc.GetConnection(appName)
	if err != nil {
		fmt.Println(err)
	}
	param, e := json.Marshal(p)
	if e != nil {
		return ""
	}
	ret, err := conn.Execute(server, "sayNice", string(param))

	if err != nil {
		fmt.Println(err)
		return ""
	}
	return ret
}
