package ch5_test

import (
	"errors"
	"fmt"
	"testing"
)

// go语言没有异常机制
// error类型实现了error接口，可以通过errors.New()快速创建错误实例

func GetFibonacci(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, errors.New("n must in [2,100]")
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestFibo(t *testing.T) {
	if ret, err := GetFibonacci(10); err == nil {
		fmt.Println(ret)
		return
	}
	fmt.Println("Params is Error")
}

// 定义错误类型，用户作为返回值，使用errors
var LessThanTwoError = errors.New("n must be more than 2")

func GetFibonacci2(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestFibo2(t *testing.T) {
	if ret, err := GetFibonacci2(1); err == nil {
		fmt.Println(ret)
		return
	}
	fmt.Println("Params is Error")
}
