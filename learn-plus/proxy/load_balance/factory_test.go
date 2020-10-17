package load_balance

import (
	"log"
	"net"
	"net/http"
	"testing"
	"time"
)

var (
	addr      = "127.0.0.1:2002"
	transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, //连接超时
			KeepAlive: 30 * time.Second, //长连接超时时间
		}).DialContext,
		MaxIdleConns:          100,              //最大空闲连接
		IdleConnTimeout:       90 * time.Second, //空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, //tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  //100-continue状态码超时时间
	}
)



func TestFactory(t *testing.T) {
	rb := LoadBanlanceFactory(LbConsistentHash)
	if err := rb.Add("http://127.0.0.1:2003/base", "10"); err != nil {
		log.Println(err)
	}
	if err := rb.Add("http://127.0.0.1:2004/base", "20"); err != nil {
		log.Println(err)
	}
	proxy := NewMultipleHostsReverseProxy(rb)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}