package main

import (
	"fmt"
	"learn-plus/Tcp/unpack"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf, err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("read from connect failed , err: %v\n", err)
			break
		}
		str := string(buf)
		fmt.Printf("receive from client, data: %v\n",str)
	}
}

func main(){
	// 1、监听端口
	listern, err := net.Listen("tcp","0.0.0.0:9999")
	if err != nil {
		fmt.Printf("listen fail,err: %v\n",err)
	}

	// 2、建立套接字连接
	for {
		conn, err := listern.Accept()
		if err != nil {
			fmt.Printf("accept fail,err : %v\n", err)
			continue
		}
		// 3、协程处理监听服务
		go process(conn)
	}
}