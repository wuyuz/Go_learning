package ch7_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {

	// 为什么这里可以，因为函数调用的时候是值复制，相当于i在传入函数时被复制了一份，所在的内存地址不一样，不存在竞争关系
	for i := 1; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Microsecond * 50)
	// 区别上面一种和下面一种的协程的输出结果？
	fmt.Println("-------Test1--------")

	// 共享变量i，多个协程共享一个变量，要达到效果可以加锁
	for i := 1; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Microsecond * 50)
}
