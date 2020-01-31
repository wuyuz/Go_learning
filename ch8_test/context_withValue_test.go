package ch8_test

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithVal(t *testing.T) {
	ProcessRequest("admin", "admin888")
}

func ProcessRequest(UserName, PassWord string) {
	// 携带数据的节点
	ctx := context.WithValue(context.Background(), "UserName", UserName)
	ctx = context.WithValue(ctx, "PassWord", PassWord)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf("处理响应 用户名:%v 密码:%v",
		ctx.Value("UserName"),
		ctx.Value("PassWord"),
	)
}

// 使用不同的自定义类型去定义同一个基本类型，在使用的时候根据类型去取值
func TestFunction_forth(t *testing.T) {
	ProcessRequest("jane", "abc123")
}

type ctxKey int

const (
	// 包装数据
	ctxUserName ctxKey = iota
	ctxPassWord
)

// 使用函数返回具体的值
func UserName(c context.Context) string {
	return c.Value(ctxUserName).(string)
}

func PassWord(c context.Context) string {
	return c.Value(ctxPassWord).(string)
}

func ProcessRequest_x(UserName, PassWord string) {
	ctx := context.WithValue(context.Background(), ctxUserName, UserName)
	ctx = context.WithValue(ctx, ctxPassWord, PassWord)
	HandleResponse(ctx)
}

func HandleResponse_x(ctx context.Context) {
	fmt.Printf(
		"处理响应 用户名:%v 密码:%v",
		UserName(ctx),
		PassWord(ctx),
	)
}
