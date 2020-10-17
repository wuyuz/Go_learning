package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
)

var addr = "127.0.0.1:2002"

func main() {
	//127.0.0.1:2002/xxx
	//127.0.0.1:2003/base/xxx
	rs1 := "http://127.0.0.1:2003/base"
	url1, err1 := url.Parse(rs1)
	if err1 != nil {
		log.Println(err1)
	}
	proxy := NewSingleHostReverseProxy(url1)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}

func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	//http://127.0.0.1:2002/dir?name=123
	//RayQuery: name=123
	//Scheme: http
	//Host: 127.0.0.1:2002
	targetQuery := target.RawQuery
	// 请求体参数修改
	director := func(req *http.Request) {
		//url_rewrite
		//127.0.0.1:2002/dir/abc ==> 127.0.0.1:2003/base/abc ??
		//127.0.0.1:2002/dir/abc ==> 127.0.0.1:2002/abc
		//127.0.0.1:2002/abc ==> 127.0.0.1:2003/base/abc
		re, _ := regexp.Compile("^/dir(.*)");
		req.URL.Path = re.ReplaceAllString(req.URL.Path, "$1")  // 正则掉所有dir后面的地址，放在新的path中

		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host

		//target.Path : /base
		//req.URL.Path : /dir
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			req.Header.Set("User-Agent", "")
		}
	}

	// 定义返回响应的函数
	modifyFunc := func(res *http.Response) error {
		// 错误码时返回
		if res.StatusCode != 200 {
			return errors.New("error statusCode")
		}
		// 正常情况下直接修改返回体
		oldPayload, err := ioutil.ReadAll(res.Body)  // 拿到返回内容
		if err != nil {
			return err
		}
		newPayLoad := []byte("hello " + string(oldPayload))  // 追加新内容
		res.Body = ioutil.NopCloser(bytes.NewBuffer(newPayLoad))  // NopCloser用一个无操作的Close方法包装r返回一个ReadCloser接口。因为res.Body是一个io.ReaderClose
		res.ContentLength = int64(len(newPayLoad))  // 修改响应长度
		res.Header.Set("Content-Length", fmt.Sprint(len(newPayLoad)))
		return nil
	}

	// 如果modifyFunc有返回error，则触犯这个错误回调
	errorHandler := func(res http.ResponseWriter, req *http.Request, err error) {
		res.Write([]byte(err.Error()))
	}
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyFunc, ErrorHandler: errorHandler}
}

func singleJoiningSlash(a, b string) string {
	// URL路径拼接
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
