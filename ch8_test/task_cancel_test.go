package ch8_test

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

// 关闭广播，会取消所有阻塞的接受者，同样会取消所有任务
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancelFunc(t *testing.T) {
	cancelChan := make(chan struct{}, 0)

	// 起5个协程，同时对一个channel作用
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, "Cancled Task!")
		}(i, cancelChan)
	}
	// 只会取消一个任务
	// cancel_1(cancelChan)

	// 会取消所有任务
	cancel_2(cancelChan)
}
