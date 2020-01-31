package ch7_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounterThreadError(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 2)
	// 发现counter输出4522，说明发生了协程不安全。也就是说有多个协程拿着同一个counter去加一了
	fmt.Println("数据不安全counter：", counter)
}

// Mutex创建数据安全环境
func TestCounterMutexSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			// Mutex总是和defer连用
			defer func() {
				// 解锁
				mut.Unlock()
			}()
			// 上锁
			mut.Lock()
			counter++
		}()
	}
	// 可能存在外面的程序执行完后终止，导致部分协程成为僵尸，所以等待
	time.Sleep(time.Second * 2)
	// 正确输出counter：5000，说明内存安全
	fmt.Println("Mutex数据安全counter：", counter)
}

// WaitGroup 类似于Python的join，表示协程/线程等待结束后才往下走，之前我们通过sleep的方式，等待协程执行完
// 但是我们并不知道协程什么时候执行完，所以我们应该使用WaitGroup来阻塞所有协程
func TestWaitGroupSafe(t *testing.T) {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		// 表示添加了一个协程计数到wg中
		wg.Add(1)
		go func() {
			defer func() {
				// 如果不在协程中解锁会导致死锁
				mu.Unlock()
			}()
			mu.Lock()
			counter++
			// 表示完成了一个协程，所以减一，类似于生产消费者栈
			wg.Done()
		}()
	}
	// 类似于join阻塞在这，等待所有完成,减少了所有协程的等待时长
	wg.Wait()
	fmt.Println("WaitGroup执行结果的counter：", counter)
}
