package ch10

import (
	"errors"
	"time"
)

// 常见的数据库链接池
type SomeObj struct {
}

type ObjPool struct {
	bufChan chan *SomeObj // 用于缓冲可用对象
}

func NewObjPool(numObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *SomeObj, numObj)
	for i := 0; i < numObj; i++ {
		objPool.bufChan <- &SomeObj{}
	}
	return &objPool
}

// 对象池中获取对象
func (p *ObjPool) GetObj(timeOut time.Duration) (*SomeObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeOut): // 超时控制
		return nil, errors.New("Time out!")
	}
}

// 释放对象到池中
func (p *ObjPool) ReleaseObj(obj *SomeObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("Overflow!")
	}
}
