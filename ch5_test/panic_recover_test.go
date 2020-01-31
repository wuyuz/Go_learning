package ch5_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

// panic 用于不可恢复的错误，它结束前会执行defer指定的内容；会输出当前调用栈信息
// os.Exit 退出时不会调用defer；不会输出调用栈信息
func TestPanicFunc(t *testing.T) {
	defer func() {
		fmt.Println("Defer functions Print out!")
	}()

	defer fmt.Println("Second Print")

	fmt.Println("Starting!")
	panic("Error") // panic(errors.New("Error of Panic"))
}

func TestOsFunc(t *testing.T) {
	defer fmt.Println("Os defer actioning?")
	fmt.Println("Os starting!")
	os.Exit(-1)
}

// recover函数 类似于Python的Exception，万能异常捕捉，可以接受到发生的报错，根据报错我们可以在
// defer中执行一些动作

func TestRecoverFunc(t *testing.T) {
	defer func() {
		// 使用defer来拦截所有报错，并打印出来，是整个函数不会抛错
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Starting...")
	panic(errors.New("Recover Error..."))
}
