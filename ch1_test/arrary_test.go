package ch1_test

import (
	"fmt"
	"testing"
)

// 切片知识
func TestScliceInit(t *testing.T) {
	fmt.Println("Test Func OK!")
	t.Log("ok") // t.Log() 如果无输出就在终端中输入go test -v xx.go
}

func TestSclice(t *testing.T) {
	var s []int
	t.Log(len(s), cap(s))
	s = append(s, 1)
	t.Log(len(s), cap(s))

	s1 := []int{1, 3, 5}
	t.Log(s1)

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 3)
	t.Log(s2)
}

//数组知识
func TestArray(t *testing.T) {
	var s [3]int
	t.Log(s)

	s1 := [4]int{1, 4, 5, 2}
	t.Log(s1)

	s2 := [...]int{2, 4, 1}
	t.Log(s2)
}

func TestArrayTravel(t *testing.T) {
	var s = [...]int{1, 32, 12, 3}
	// 方式一
	for i := 0; i < len(s); i++ {
		t.Log(s[i])
	}

	// 方式二
	for index, val := range s {
		t.Log(index, val)
	}
}
