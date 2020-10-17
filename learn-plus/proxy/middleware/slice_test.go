package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"
)

var addr = "127.0.0.1:2002"

func TestSliceRouter(t *testing.T) {
	reverseProxy := func(c *SliceRouterContext) http.Handler {
		rs1 := "http://127.0.0.1:2003/base"
		url1, err1 := url.Parse(rs1)
		if err1 != nil {
			log.Println(err1)
		}

		rs2 := "http://127.0.0.1:2004/base"
		url2, err2 := url.Parse(rs2)
		if err2 != nil {
			log.Println(err2)
		}

		urls := []*url.URL{url1, url2}
		return NewMultipleHostsReverseProxy(c,urls)
	}
	log.Println("Starting httpserver at " + addr)

	//初始化方法数组路由器
	sliceRouter := NewSliceRouter()

	//中间件可充当业务逻辑代码
	sliceRouter.Group("/base").Use(TraceLogSliceMW(), func(c *SliceRouterContext) {
		c.Rw.Write([]byte("test func"))
	})

	//请求到反向代理
	sliceRouter.Group("/").Use(TraceLogSliceMW(), func(c *SliceRouterContext) {
		fmt.Println("reverseProxy")
		reverseProxy(c).ServeHTTP(c.Rw, c.Req)
	})
	routerHandler := NewSliceRouterHandler(nil, sliceRouter)
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}
