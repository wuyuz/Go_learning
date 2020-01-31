package ch9_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Microsecond * 50)
	return fmt.Sprintf("The result is from task: %d", id)
}

func FirstResponse() string {
	numRunner := 10
	// ch := make(chan string)
	ch := make(chan string, numRunner)
	// 同时起10个协程，并发执行任务，凡是有一个结果写进管道中
	for i := 0; i < numRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	// 由于上面的十个协程凡是有一个写了数据进管道中，就会即可return，终止函数
	return <-ch
}

// 任意一个任务完成即完成；
func TestFirstFunc(t *testing.T) {
	// 打印运行前的协程数：2
	t.Log("Before: ", runtime.NumGoroutine())
	// 每次返回的结果都不一样，这和协程的调度机制是有关的
	t.Log(FirstResponse())
	//  睡眠一段时间保证所有协程完成
	time.Sleep(time.Microsecond * 500)
	// 打印运行后的协程数：11， 按理说应该有12条的，但是有一条成功结束，也就是说这里产生了阻塞的9条僵尸协程，如果过多的协程，会
	// 造成out of memory！是相当危险的。为了防止这种情况：我们一定要使用带buffer的管道，指定长度
	t.Log("After: ", runtime.NumGoroutine())
}

func AllResultFunc() string {
	number := 10
	ch := make(chan string, 10)
	for i := 0; i < number; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	AllResult := ""
	for j := 0; j < number; j++ {
		AllResult += <-ch + "\n"
	}
	return AllResult
}

// 获得所有的结果后返回;当然我们也可以使用sync.WaitGroup来完成
func TestAllResult(t *testing.T) {
	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(AllResultFunc())
	t.Log("After: ", runtime.NumGoroutine())
}
