package ch9_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

// 使用全局的实例对象
var SingletonInstance *Singleton

// 用于创建单例模式
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj！")
		SingletonInstance = new(Singleton)
	})
	return SingletonInstance
}

// 单例模式
func TestSingleton(t *testing.T) {
	// 使用sync的WaitGroup来等待所有协程结束
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("Obj 对象：%x \n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
