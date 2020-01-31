package ch11_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// sync.poo区别于池的概念，更应该叫做：cache_pool,因为它会受GC影响，缓存的数据，首先回去processor的私有对象中获取，没有的话
// 会到processor的共享数据中获取，所有的池中都没有，就会触发自定义的New函数； 需要注意，共享池就是起的协程能拿到的数据

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a New obj ")
			// 由于返回的是interface，所以最好做断言
			return 100
		},
	}

	// 断言的返回值是一个 bool
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() // 调用了GC后，清除所有缓存对象，所有put的私有变量都会失效，当下次获取的时候就会触发New函数
	if v1, err := pool.Get().(int); err == true {
		fmt.Println("Success get msg: ", v1)
	} else {
		fmt.Println("Get wrong: ", err)
	}
}

func TestMultiSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a New obj ")
			// 由于返回的是interface，所以最好做断言
			return 10
		},
	}
	pool.Put(2)
	pool.Put(2)
	pool.Put(2)

	var wg sync.WaitGroup
	// 起十个协程，他们会共享3个2没有的会触发New
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			t.Log(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
	time.Sleep(time.Second * 1)
}
