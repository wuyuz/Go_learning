package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr     *cronexpr.Expression
		err      error
		now      time.Time
		nextTime time.Time
	)

	// 每5秒执行一次
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	// 当前时间
	now = time.Now()
	// 下次调度时间
	nextTime = expr.Next(now)
	//fmt.Println(now,nextTime)

	//等待这个定时器超时执行
	time.AfterFunc(nextTime.Sub(now), func() {  // 等待现在和下一个时间的间隔后执行
		fmt.Println("nextTime时间到达了：", nextTime)
	})
	time.Sleep(6 * time.Second)

}
