package main

import (
	"fmt"
	"os/exec"
)

func main(){
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)

	// 生成CMD
	cmd = exec.Command("/bin/bash","-c","sleep 5; ls -l")
	// 执行cmd命令并获取输出
	if output, err = cmd.CombinedOutput();err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))
}