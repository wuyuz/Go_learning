package ch7_test

import (
	"fmt"
	"runtime"
	"testing"
)

// 这个函数的作用是让当前goroutine让出CPU，好让其它的goroutine获得执行的机会。
// 同时，当前的goroutine也会在未来的某个时间点继续运行。

func showNumber(i int) {
	fmt.Println("show Num: ", i)
}

// 未切换协程
func TestOneFunction(t *testing.T) {
	for i := 0; i < 8; i++ {
		// showNumber(i)      // 能顺利打印出8个数字
		go showNumber(i) // 不能打印全number信息，因为主进程退出
	}
}

func TestTwoFunction(t *testing.T) {
	for i := 0; i < 10; i++ {
		go showNumber(i)
	}
	// 由于这里，表示可以切换到go起的协程中进行执行，表示让出cpu执行权给协程
	runtime.Gosched()
	fmt.Println("主线程退出")
}
