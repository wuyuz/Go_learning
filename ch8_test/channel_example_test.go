package ch8_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 说明下面函数的结果，下面的编程方式是否安全？
func TestFunction(t *testing.T) {
	number := 10
	// ch := make(chan int)
	ch := make(chan int, 10) // 带buffer的channel避免僵尸协程
	fmt.Println("Before: ", runtime.NumGoroutine())
	for i := 0; i < number; i++ {
		go func(i int) {
			fmt.Printf("协程：%d，readly！\n", i)
			ch <- i
			fmt.Printf("协程：%d，end！\n", i)
		}(i)
	}
	fmt.Printf("first result of channel: %d \n", <-ch)
	// 保证协程完成运行
	time.Sleep(time.Second * 1)
	fmt.Printf("Second result of channel: %d \n", <-ch)
	fmt.Println("After: ", runtime.NumGoroutine())
}

// Before:  2
// 协程：9，readly！
// 协程：9，end！
// first result of channel: 9
// 协程：2，readly！
// 协程：3，readly！
// 协程：4，readly！
// 协程：5，readly！
// 协程：6，readly！
// 协程：7，readly！
// 协程：8，readly！
// 协程：1，readly！
// 协程：0，readly！
// Second result of channel: 2
// After:  11
// --- PASS: TestFunction (1.01s)
// 协程：2，end！

// 又上我们可以得到一下结论：默认chan的容量为1，仅仅支持一读一写，否则就会阻塞；在上面的程序中，会制造8个僵尸协程，
// 因为总有几个被阻塞，且无法退出！我们可以使用带buffer的channel，来防止这种情况，让协程都有可写的channel，这样就避免了阻塞
