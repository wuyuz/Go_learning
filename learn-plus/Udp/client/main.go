package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听服务器
	conn, err :=  net.DialUDP("udp",nil, &net.UDPAddr{
		IP:   net.IPv4(127,0,0,1),
		Port: 9909,
	})

	if err != nil {
		fmt.Println("listen failed")
		return
	}
	for i := 0; i <100; i++ {
		_,err = conn.Write([]byte("hello server"))
		if err != nil {
			fmt.Printf("send data failed :%v", err)
			return
		}

		result := make([]byte,1024)
		n, remoteAddr, err := conn.ReadFromUDP(result)
		if err != nil {
			fmt.Printf("receive data failed %v",err)
			return
		}
		fmt.Printf("recevice data: %v, addr: %v,\n", string(result[:n]),remoteAddr)
	}
}
