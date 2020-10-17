package main

import (
	"log"
	"net/http"
	"time"
)

var (
	Addr = ":1212"
)

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Write([]byte("bye bye"))
}

func main() {
	// 创建路由器
	mux := http.NewServeMux()
	// 设置路由规则
	mux.HandleFunc("/bye",sayBye)

	// 创建服务器
	server := &http.Server{
		Addr: Addr,
		WriteTimeout: time.Second*3,
		Handler: mux,
	}

	// 监听端口并提供服务
	log.Println("Starting httpServer at ", Addr)
	log.Fatal(server.ListenAndServe())
}