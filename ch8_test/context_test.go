package ch8_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// context是Go中广泛使用的程序包，由Google官方开发，在1.7版本引入。它用来简化在多个go routine传递上下文数据、
// (手动/超时)中止routine树等操作，比如，官方http包使用context传递请求的上下文数据，gRpc使用context来终止某个请求产生
// 的routine树。每个Context应该视为只读的，通过WithCancel、WithDeadline、WithTimeout和WithValue函数可以基于现有的
// 一个Context（称为父Context）派生出一个新的Context（称为子Context）。其中WithCancel、WithDeadline和WithTimeout
// 函数除了返回一个派生的Context以外，还会返回一个与之关联CancelFunc类型的函数，用于关闭Context。通过调用CancelFunc来关
// 闭关联的Context时，基于该Context所派生的Context也都会被关闭，并且会将自己从父Context中移除，停止和它相关的timer。

// context是go语言官方用于处理取消/传递消息基于关联任务的一种实现模块
// 1、 根context：通过context.Background()创建
// 2、 子context：通过context.WithCancel(partentContext) 创建，并返回一个context对象+一个关闭它的函数
// 3、 关闭parent节点：会连带关闭其下的子节点，关闭消息会出给ctx.Done()

func isCancelle(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	// 显然这里使用Background()创建的是根节点，返回context对象，和关闭根节点的函数
	ctx, cancle := context.WithCancel(context.Background())

	// 起5个协程，同时对一个channel作用
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelle(ctx) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, "Cancled Task!")
		}(i, ctx)
	}
	// 取消
	cancle()
	time.Sleep(time.Microsecond * 50)
}
