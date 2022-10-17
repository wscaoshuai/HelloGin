package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter,r *http.Request){
	//定义模板
	t := template.New("f.tmpl")//创建一个名为f的模板对象
	//解析模板
	_,err := t.ParseFiles("./f.tmpl")
	if err != nil{
		fmt.Printf("tmpl error%v\n",err)
		return
	}
	name := "铁傀儡"
	//渲染模板
	t.Execute(w,name)
}

func main() {
	http.HandleFunc("/",f1)
	err := http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Printf("http server error:%v",err)
		return
	}
}