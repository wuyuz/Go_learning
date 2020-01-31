package service

import (
	"errors"
	"fmt"
)

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

func init() {
	fmt.Println("init function one")
}

func Square(n int) int {
	return n * n
}

// 在main函数执行之前，所有导入包的init函数都会按照导入顺序执行
// 每个包可以有多个init函数
// 引用同一包的都个函数，也只会执行一次init函数

func init() {
	fmt.Println("init function two")
}
