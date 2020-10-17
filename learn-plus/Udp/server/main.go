package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听服务器
	listen, err :=  net.ListenUDP("udp",&net.UDPAddr{
		IP:   net.IPv4(0,0,0,0),
		Port: 9909,
	})

	if err != nil {
		fmt.Println("listen failed")
		return
	}

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read failed from addr: %s",addr)
			break
		}

		go func() {
			fmt.Printf("read failed from addr: %s ,data:%v, count:%v\n",addr, string(data[:n]),n)
			_, err := listen.WriteToUDP([]byte("received sucess"),addr)
			if err != nil {
				fmt.Printf("write failed : %v",err)
			}
		}()
	}
}