package ch8_test

import (
	"fmt"
	"testing"
)

//  斐波那契函数
func Fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x: // 只要quit没有数据，就每次往里面写x
			x, y = y, x+y
		case <-quit: // 处于阻塞，只要监听到数据就退出
			fmt.Println("quit select!")
			// fmt.Println(<-c)
			return
		}
	}
}

func TestFibonacciFunc(t *testing.T) {
	// c := make(chan int, 10)
	c := make(chan int, 1) // 使用带buffer的select
	quit := make(chan int)
	go func() {
		for i := 1; i < 7; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	Fibonacci(c, quit)
}

// Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。就是用来监听阻塞的的多路选择机制，那条路可走就走
// select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作，大致的结构如下：

//select {
//　　case <-chan1:
// 如果chan1成功读到数据，则进行该case处理语句
//　 case chan2 <- 1:
// 如果成功向chan2写入数据，则进行该case处理语句
//　　default:
// 如果上面都没有成功，则进入default处理流程
//}

// 在一个select语句中，Go语言会按顺序从头至尾评估每一个发送和接收的语句。
// 如果其中的任意一语句可以继续执行(即没有被阻塞)，那么就从那些可以执行的语句中任意选择一条来使用。
// 如果没有任意一条语句可以执行(即所有的通道都被阻塞)，那么有两种可能的情况：
//    1、如果给出了default语句，那么就会执行default语句，同时程序的执行会从select语句后的语句中恢复。
//    2、如果没有default语句，那么select语句将被阻塞，直到至少有一个通信可以进行下去。
