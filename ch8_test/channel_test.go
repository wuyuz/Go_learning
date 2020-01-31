package ch8_test

import (
	"fmt"
	"testing"
	"time"
)

// go语言中使用两种channel：1、一个channel的两端保持一致性即：A比较在channel中将消息交给B才能走，否则阻塞；
// 2、A只管往channel中放东西，B有时间就去取
func service() string {
	time.Sleep(time.Microsecond * 50)
	return "Done Service"
}

func OtherTask() {
	fmt.Println("working on something else.")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("other task Done.")
}

// 此为同步操作，耗时是执行程序耗时之和，1.7s多
func TestService(t *testing.T) {
	fmt.Println(service())
	OtherTask()
}

// 上面同步有个耗时的操作，因为不管快慢都要sleep 1s钟，使用channel，有了就立即返回会大大提升效率
// 异步调用
func AsyncService() chan string {
	// retCh := make(chan string)  // 这种写法是不够高效的，由于它会等到channel中的数据被取出后才会往下执行，也就是没有使用buffer的channel

	retCh := make(chan string, 1) // 由于此channel长度为1（buffer的channel），写进一个数据后就里面结束，高效

	// 需要起一个协程，不然还是同步，这个协程执行完后会将结果立即给chaneel
	go func() {
		ret := service()
		fmt.Println("return resulted.")
		retCh <- ret
		fmt.Println("exited service.")
	}()
	return retCh
}

// 耗时明显缩短0.6
func TestAysncFunc(t *testing.T) {
	retch := AsyncService()
	// 由于AsyncService是一个异步的协程，所以不会阻塞
	OtherTask()          // 马上执行
	fmt.Println(<-retch) // 太快，会阻塞在此
	time.Sleep(time.Microsecond * 50)
}

func TestTimeOutAysncFunc(t *testing.T) {
	// 由于AsyncService是一个异步的协程，所以不会阻塞
	OtherTask() // 马上执行

	// 多路选择+时间超时返回异常，由于channel和时间阻塞都是io操作，所以配合select多路选择很实用
	select {
	case ret := <-AsyncService():
		fmt.Println(ret) // 太快，会阻塞在此
	case <-time.After(time.Microsecond * 50):
		fmt.Println("Time out 50ms")
		// t.Error("time out")
	}

	time.Sleep(time.Microsecond * 50)
}
