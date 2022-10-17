package main

import (
	"fmt"
	"html/template"
	"net/http"
)


func sayhello(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("ParseFiles has error%v",err)
		return
	}

	//渲染模板
	err = t.Execute(w,"铁傀儡")
	if err != nil {
		fmt.Printf("rander template error",err)
		return
	}

}

func main() {
	http.HandleFunc("/",sayhello)
	err :=http.ListenAndServe(":8086",nil)
	if err != nil{
		fmt.Println("HttpServer has Error",err)
		return
	}
}
