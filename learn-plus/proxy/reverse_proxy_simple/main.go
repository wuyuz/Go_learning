package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var add = "127.0.0.1:2002"

func main() {
	// 需要代理的下游服务url
	rs1 := "http://127.0.0.1:2003/base"  // 实现了重写url的规则
	//127.0.0.1:2002/xxx
	//127.0.0.1:2003/base/xxx
	url1,err1 := url.Parse(rs1)  // 解析url
	if err1 != nil{
		log.Println(err1)
	}
	proxy := httputil.NewSingleHostReverseProxy(url1) // 传递个url进去后，实现了一个direct，然后注册到ReverseProxy并返回
	log.Println("Starting httpServer at: "+add)
	log.Fatal(http.ListenAndServe(add,proxy)) // proxy中实现了Handler接口中的ServeHTTP
}