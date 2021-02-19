package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {
	// 需要一个调度协程，定检查所有的cron任务，谁要过期了就执行
	var (
		cronJob *CronJob
		expr *cronexpr.Expression
		now time.Time
		scheduleTable map[string]*CronJob
	)

	scheduleTable= make(map[string]*CronJob)

	// 当前时间
	now = time.Now()
	// 创建第一个job任务
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:expr,
		nextTime: expr.Next(now),
	}
	// 注册任务
	scheduleTable["job1"]=cronJob

	// 创建第二个job任务
	expr = cronexpr.MustParse("*/7 * * * * * *")
	cronJob = &CronJob{
		expr:expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job2"]=cronJob

	//启动一个调度协程
	go func() {
		var(
			now time.Time
			cronJobName string
			cronJob *CronJob
		)

		// 定时检查下个任务
		for {
			now = time.Now()
			for cronJobName,cronJob = range scheduleTable{
				// 是否过期
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now){
					//启动一个协程
					go func(cronJobName string) {
						fmt.Println("执行：",cronJobName)
					}(cronJobName)

					// 计算下次调度时间
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println("下次执行时间：",cronJob.nextTime)
				}
			}

			// 阻塞一段时间,100毫秒，原理是这个Timer中有个channel，如果到期了会从这个channel中读出数据
			select {
				case <- time.NewTimer(100*time.Millisecond).C:
			}
		}

	}()


	time.Sleep(100*time.Second) //避免主协程退出

}
