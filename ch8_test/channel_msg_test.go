package ch8_test

import (
	"fmt"
	"sync"
	"testing"
)

// 生产者函数
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// 如果生产者生产商品多余消费者，那么消费者会阻塞
		for i := 1; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()
}

// 消费者函数
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 1; i < 10; i++ {
			data := <-ch
			fmt.Println("消费：", data)
		}
		wg.Done()
	}()
}

func TestChannelFunc(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	// 需要注意，在使用sync.WaitGroup的时候，一定要进行指针传递，否则值传递无法实现wg中的lock
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}

// 问题： 上面的代码存在一个问题，那就是只有一个reciver，当多个reciver存在的时候，生产者生产完数据后，被消费者消费掉，总会存在
// 其他消费者不知道生产者生产完数据，或者数据取完了的情况，一般我们的解决方案是将生产者完成生产后以特定的token，如：-1表示取完了
// 要求关闭一个reciver，那么生产者要根据多少个reciver生产多少个-1去关闭，这显得很麻烦！

// 为了解决上面的问题：我们使用了channel的关闭广播机制
//   1、 向关闭的channel发送数据会抛panic
//   2、 v, ok <- ch; ok为true表示正确获取数据；为false表示通道关闭，所有channel的接受者，都会在通道关闭时立即从
//  阻塞状态返回上诉ok为false，且返回chan的默认0值，表示通道关闭，退出信号。

// 添加关闭状态判断的接受者
func dataReceiverClose(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			// 凡是关闭通道，ok就会返回false，从而break'
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

// 关闭channel的生产者
func dataProducerClose(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 生产完后就关闭channel，会广播所有的订阅者，如果不关闭则消费者会一致阻塞
		close(ch)
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	dataProducerClose(ch, &wg)

	// 多个接受者，生产者关闭后，自动退出接受者
	wg.Add(1)
	dataReceiverClose(ch, &wg)
	wg.Add(1)
	dataReceiverClose(ch, &wg)
	wg.Wait()
}
