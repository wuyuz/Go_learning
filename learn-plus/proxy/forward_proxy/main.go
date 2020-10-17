package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Pxy struct{}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n",req.Method,req.Host,req.RemoteAddr)
	transport := http.DefaultTransport   // 数据连接池
	// step1, 浅拷贝对象，然后就再新增属性数据，因为浅拷贝可以保存新增属性
	outReq := new(http.Request)
	*outReq = *req

	if clienIP,_,err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior,ok := outReq.Header["X-Forwarded-For"];ok { // X-Forwarded-For 是一个代理Ip的链表
			clienIP = strings.Join(prior,", ") + ", "+clienIP
		}
		outReq.Header.Set("X-Forwarded-For",clienIP)
	}

	// step2, 请求下游
	res, err := transport.RoundTrip(outReq) // 获取到下游的响应数据
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	// step3, 下游请求内容返回给上游
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key,v)  // 头写入
		}

	}

	//rw.Header().Add("xxx","sss")  // 可以给每个请求进行修改
	rw.WriteHeader(res.StatusCode)
	io.Copy(rw,res.Body)  // 数据写入
	res.Body.Close()
}

func main() {
	fmt.Println("Serve on :8080")
	http.Handle("/",&Pxy{})  // Handle中的第二个参数必须实现ServeHTTP
	http.ListenAndServe("0.0.0.0:8080",nil)
}