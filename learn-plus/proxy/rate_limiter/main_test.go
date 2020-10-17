package rate_limiter

import (
	"learn-plus/proxy/middleware"
	"log"
	"net/http"
	"net/url"
	"testing"
)

var addr = "127.0.0.1:2002"

// 熔断方案
func TestRater(t *testing.T) {
	coreFunc := func(c *middleware.SliceRouterContext) http.Handler {  // 创建一个核心方法
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
		return middleware.NewMultipleHostsReverseProxy(c, urls)
	}
	log.Println("Starting httpserver at " + addr)

	sliceRouter := middleware.NewSliceRouter()  // 创建一个我们自己封装的路由
	sliceRouter.Group("/").Use(RateLimiter())
	routerHandler := middleware.NewSliceRouterHandler(coreFunc, sliceRouter) // 传入的是核心方法handler，和一个路由
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}