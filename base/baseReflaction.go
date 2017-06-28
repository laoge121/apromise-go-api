package base

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	age  int
}

type Manage struct {
	user  User
	title string
}

func (u User) Hellow() {
	fmt.Println("hellow word!")
}

func Info(o interface{}) {

	t := reflect.TypeOf(o)
	fmt.Println("type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("fields", v)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		f := t.Method(i)
		fmt.Println(f.Name, f.Type)
	}
}
