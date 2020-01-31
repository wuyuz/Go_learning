package ch10

import (
	"fmt"
	"testing"
	"time"
)

// 在运行代码的时候，如果引用了其他包，需要运行：go test -v *.go
func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("get Obj: %T \n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}
