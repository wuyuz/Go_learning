package ch7_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"
)

// 读写锁是针对于读写操作的互斥锁,基本遵循两大原则：
// 1、可以随便读。多个goroutin同时读。因为读和读不互斥；区别于Mutex锁，读锁被释放了，而Mutex对每一步都进行了lock，读些锁提高了效率
// 2、写的时候，啥都不能干。不能读，也不能写。写锁于其他操作都互斥

// 问题：如果一个写的操作刚执行完了第一个指令，时间片换给另一个读的协程，这就会读到一个错误的数据。

var rm sync.RWMutex

func TestRMutexFunc(t *testing.T) {
	fmt.Println("TestFunction Starting")
	fmt.Println("Test function addr: ", unsafe.Pointer(&rm))
	// 读锁没有限制
	go read(1)
	go read(2)
	time.Sleep(time.Second * 4)
}

// 读锁
func read(i int) {
	fmt.Println(i, "read start")
	// rm 使用的是全局变量
	fmt.Println("read function rm's add: ", unsafe.Pointer(&rm))
	rm.RLock()
	fmt.Println(i, "reading")
	time.Sleep(1 * time.Second)
	rm.RUnlock()
	fmt.Println(i, "read end")
}

// 写锁
func write(i int) {
	println(i, "write start")
	rm.Lock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	rm.Unlock()
	println(i, "write end")
}

func TestWriteFunc(t *testing.T) {
	// 写操作锁
	go write(1)
	go read(2)
	go write(3)
	time.Sleep(5 * time.Second)
}

// 3 write end结束之后，2才能reading;  2 read end结束之后，1 才能writing
