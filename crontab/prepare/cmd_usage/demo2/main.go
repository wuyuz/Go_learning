package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err error
	output []byte
}

func main()  {
	// 起一个协程执行一个等待3秒的程序，最外层在等待1秒时停止协程的命令
	var (
		ctx context.Context
		ctxFunc context.CancelFunc
		cmd *exec.Cmd
		resultChan chan result
		res result
	)

	// ctx 中又一个chan，ctxFunc可以关闭这个chan，如果调用ctx的函数监听了这个上下文的channel(select { case <- ctx.Done})，如果关闭就结束
	ctx,ctxFunc = context.WithCancel(context.TODO())

	// 初始化一个结果队列
	resultChan = make(chan result,10)

	go func() {
		var (
			output []byte
			err error
		)
		cmd = exec.CommandContext(ctx,"/bin/bash","-c","sleep 3; ls -l")

		output, err = cmd.CombinedOutput()

		resultChan <- result{
			err,
			output,
		}
	}()

	time.Sleep(4*time.Second)
	// 取消上下文
	ctxFunc()

	res = <- resultChan
	fmt.Println(res.err, string(res.output))
}
