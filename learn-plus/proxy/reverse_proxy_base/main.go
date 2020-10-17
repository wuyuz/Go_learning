package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
)

var (
	proxy_addr = "http://127.0.0.1:2003"
	port = "2002"
)

func handler(w http.ResponseWriter,r *http.Request) {
	// step1 解析代理地址，并更改请求体的协议和主体
	proxy, err := url.Parse(proxy_addr)
	r.URL.Scheme = proxy.Scheme // 协议
	r.URL.Host = proxy.Host

	// step2 请求下游服务
	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(r)

	if err != nil {
		log.Println(err)
		return
	}

	// step3 把下游请求内容返回给上游
	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k,v)
		}
	}

	defer resp.Body.Close()
	bufio.NewReader(resp.Body).WriteTo(w)  // 将拿到的下游响应写入
}

func main() {
	http.HandleFunc("/",handler)
	log.Println("Start serving on port: ",port)
	err := http.ListenAndServe(":"+port,nil)
	if err != nil {
		log.Fatal(err)
	}
}
