package ch2_test

import (
	"fmt"
	"testing"
	"time"
)

// 闭包函数,打印内部函数执行时间
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("spent time: ", time.Since(start).Seconds())
		return ret
	}
}

// 需要检测的执行函数
func slowFunc(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

func TestFn(t *testing.T) {
	obj := timeSpent(slowFunc)
	ret := obj(11)
	fmt.Println("测试闭包函数完成：", ret)
}
