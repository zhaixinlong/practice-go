// main.go

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse file failed, err", err)
		return
	}
	//渲染模板
	//u1 := User{
	//	Name:   "eryajf",
	//	Gender: "男",
	//	Age:    10,
	//}
	//通过结构体传递的话，字段名称首字母必须大写才能访问到
	//而通过map的形式，则不需要首字母大写，因为是直接通过key访问的
	m1 := map[string]interface{}{
		"Name":   "小明",
		"Gender": "男",
		"Age":    20,
	}
	t.Execute(w, m1) // 此处传递结构体，或者map都是可以的
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Http server start failed, err", err)
		return
	}
}
