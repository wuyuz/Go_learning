package ch1_test

import "testing"

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
