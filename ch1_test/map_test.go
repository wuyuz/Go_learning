package ch1_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	m1 := map[int]int{1: 2, 3: 4}
	t.Log(m1)

	m2 := map[int]int{}
	m2[3] = 98
	t.Log(m2)

	// 此方法初始化是为了一次性创建固定长度的map，省去了重复创建的开销
	m3 := make(map[int]int, 3)
	m3[2] = 12
	t.Log(m3)

	if v, ok := m3[2]; ok {
		t.Log("获取成功：", v)
	} else {
		t.Log("获取失败")
	}
}

// 自带的map数据类型读是数据安全的，但是写是不安全的，当读写竞争时会抛错，所以使用sync.Map.
// sync.Map 有以下特性：
// 1、无须初始化，直接声明即可。
// 2、sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
// 3、使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，
// 返回 true，终止迭代遍历时，返回 false。

func TestSyncMap(t *testing.T) {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

}
