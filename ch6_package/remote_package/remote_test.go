package remote_package

import (
	"fmt"
	"testing"

	cm "github.com/easierway/concurrent_map"
)

// 传统的map不是线程安全的
func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key1"), 10)
	if v, err := m.Get(cm.StrKey("key")); err == true {
		fmt.Println("获得值： ", v)
	} else {
		fmt.Println("未获得值！")
	}
}
