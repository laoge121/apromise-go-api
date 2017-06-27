package base

import (
	"fmt"
)

type Usb interface {
	Names() string
	Connect()
}

type Connecter interface {
	Connect()
}

type TvConnector struct {
	Name string
}

func (tv TvConnector) Connect() {
	fmt.Println(tv.Name, "连接接入！")
}

type PhoneConnector struct {
	Name string
}

func (ph PhoneConnector) Names() string {
	return ph.Name
}

func (ph PhoneConnector) Connect() {
	fmt.Println("connected", ph.Name)
}

func DisConnect(sb Usb) {
	if ph, ok := sb.(PhoneConnector); ok {
		fmt.Println(ph.Names(), "断开连接！")
	}
}

func DisConnectInf(sub interface{}) {

	switch v := sub.(type) {
	case PhoneConnector:
		fmt.Println(v.Names(), "断开连接！")
	default:
		fmt.Println(">>>>>>>")
	}
}
