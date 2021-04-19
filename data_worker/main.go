package main

import (
	"data_worker/app"
	"data_worker/common"
	. "data_worker/worker"
	"fmt"
	"time"
)

func main() {
	var (
		err error
	)
	app.Log = common.SetupZeroLog()
	app.Log.Info().Msg("[+] data worker starting...")
	err = Service()
	if err != nil {
		fmt.Println("[+] error: %s", err)
		return
	}

	time.Sleep(10 * time.Second)
}

func Service() error {
	var (
		err error
	)
	// 加载环境变量
	common.LoadEnv()
	app.Client, err = common.GetMongoClient()
	if err != nil {
		return err
	}

	// timer worker running
	go TimerWork()
	return nil
}
