package ch8_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func SayHello() {
	fmt.Println("hello from Function: SayHello")
}

// 此方法不会输出协程结果，因为没有时间给它输出
func TestHello(t *testing.T) {
	go func() { fmt.Println("hello from Test-goroutenie") }()
	go SayHello()
	fmt.Println("Hello from TestMain")
}

// 通道：如果要将消息/错误/数据从一个goroutine传递到另一个goroutine，那么就要使用channel来传递消息
func printHello(ch chan int) {
	fmt.Println("hello from Function: SayHello")
	ch <- 2
}

func TestChannel(t *testing.T) {
	ch := make(chan int)
	go func() {
		fmt.Println("Hello inline")
		//send a value on channel
		ch <- 1
	}()
	go printHello(ch)
	fmt.Println("Hello from main")
	i := <-ch
	fmt.Println("Recieved ", i)
	// 会阻塞在此
	fmt.Println("Recivered: ", <-ch)
}

// 上下文处理举例
var (
	logg *log.Logger
)

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logg.Println(`下班!`)
			return
		default:
			// 一般情况下上班
			logg.Println(`上班!`)
			time.Sleep(1 * time.Second)
		}
	}
}

func TestFunction_one(t *testing.T) {
	logg = log.New(os.Stdout, "", log.Ltime)
	//不知道这个context应该场景来干嘛就用:context.TODO()
	ctx, cancel := context.WithCancel(context.TODO())
	//5秒后结束,这个context
	time.AfterFunc(5*time.Second, func() {
		cancel()
	})
	go work(ctx)
	time.Sleep(10 * time.Second)
	logg.Println(`无脑发呆中!`)
}

func work_two(ctx context.Context, ch chan bool) {
	for {
		select {
		case <-ctx.Done():
			logg.Println(`下班!`)
			ch <- true
			return
		default:
			logg.Println(`上班!`)
			time.Sleep(2 * time.Second)
		}
	}
}

// context.WithDeadline()和context.WithTimeout():
// 返回Context和取消函数用来取消Context(这个取消函数会根据设置的时间自动取消)
func TestFunction_two(t *testing.T) {
	ch := make(chan bool)
	logg = log.New(os.Stdout, "", log.Ltime)
	//
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	go work_two(ctx, ch)
	// 使用Deadline函数会在固定的时间取消创建的对象，如果超时，这个函数会打印三次上班后下班
	time.Sleep(10 * time.Second)
	//取消函数：当cancel被调用时,context.WithDeadline设置的时间超过了，关闭ctx.Done通道。
	cancel()
	// 这个chan是为了保证子的goroutine执行完,当然也可以不用chan用time.Sleep停止几秒
	<-ch
	logg.Println(`无脑发呆中!`)
}

func work_three(ctx context.Context, ch chan bool) {
	for {
		select {
		case <-ctx.Done():
			logg.Println(`下班!`)
			ch <- true
			return
		default:
			logg.Println(`上班!`)
			time.Sleep(2 * time.Second)
		}
	}
}

func TestFunction_three(t *testing.T) {
	ch := make(chan bool)
	logg = log.New(os.Stdout, "", log.Ltime)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go work_three(ctx, ch)
	time.Sleep(10 * time.Second)
	//取消函数：当cancel被调用时,context.WithTimeout设置的时间超过后,关闭ctx.Done通道；
	cancel()
	// 这个chan是为了保证子的goroutine执行完,当然也可以不用chan用time.Sleep停止几秒
	<-ch
	logg.Println(`无脑发呆中!`)
}
