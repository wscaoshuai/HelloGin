package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age int
	Gender string
}

func sayhello(w http.ResponseWriter,r *http.Request ){
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parseFiles error", err)
		return
	}
	//渲染模板
	u1 := User{
		Name:   "铁傀儡",
		Age:    20,
		Gender: "男",
	}
	m1 := map[string]interface{}{
		"name":   "铁傀儡",
		"age":    20,
		"gender": "男",
	}
	t.Execute(w,map[string]interface{}{
		"m1" : m1,
		"u1" : u1,
	})
}

func main() {
	http.HandleFunc("/",sayhello)
	err := http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Println("http server error:",err)
		return
	}

}
