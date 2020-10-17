package ch4_test

import (
	"fmt"
	"testing"
)

// 该函数定义一个空接口参数，不受类型的限制，根据传入的参数进行断言，
func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer: ", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("String: ", s)
		return
	}
	fmt.Println("Unkown Type")
}

// 空接口常配合switch使用，简化操作
func DoSomething2(p interface{}) {
	fmt.Println(p)
	switch v := p.(type) {
	case int:
		fmt.Println("Integer: ", v)
	case string:
		fmt.Println("String: ", v)
	default:
		fmt.Println("Unkown Type")
	}
}

func TestEmpty(t *testing.T) {
	DoSomething(10)
	DoSomething2("10")
}
