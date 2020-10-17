package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	// 创建连接池
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:       30*time.Second,  // 连接超时
			KeepAlive:     30*time.Second,  // 长连接超时时间
		}).DialContext,
		MaxIdleConns:100,    // 最大空闲连接
		IdleConnTimeout:90*time.Second,  //  空闲超时使劲啊
		TLSHandshakeTimeout:10*time.Second,  //  tls握手超时时间
		ExpectContinueTimeout:1*time.Second,  //  100-continue状态吗超时时间
	}
	//  创建客户端
	client := &http.Client{
		Timeout: time.Second * 30,
		Transport:transport,
	}

	// 请求数据
	resp, err := client.Get("http://127.0.0.1:1212/bye")
	if err !=  nil {
		panic(err)
	}
	defer resp.Body.Close()  // 不释放会独占内存
	// 读取数据
	bas, err := ioutil.ReadAll(resp.Body)
	if err !=  nil {
		panic(err)
	}
	fmt.Println(string(bas))
}