package main

import (
	"learn-plus/proxy/load_balance"
	"learn-plus/proxy/middleware"
	"log"
	"net/http"
)

var (
	addr = "127.0.0.1:2002"
)

func main() {
	rb := load_balance.LoadBanlanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("http://127.0.0.1:2003", "50")
	proxy := load_balance.NewLoadBalanceReverseProxy(&middleware.SliceRouterContext{}, rb)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
