package client
import (
	"testing"
	"ch6_package/service"  // 指定到包名位置
)


func TestTryFunc(t *testing.T){
	t.Log(service.GetFibonacci(10))
}