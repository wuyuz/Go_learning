package main

import (
	"fmt"
	"learn-plus/Tcp/unpack"
	"net"
)

func main() {
	// 1、 连接服务器
	conn, err := net.Dial("tcp", "0.0.0.0:9999")
	defer conn.Close()
	if err != nil {
		fmt.Println("connect failed")
		return
	}

	unpack.Encode(conn, "hello, world 000 !!!")
}