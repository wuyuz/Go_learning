package ch12_test

import (
	"fmt"
	"reflect"
	"testing"
)

// reflect.TypeOf() 返回类型
// reflect.ValueOf() 返回值
// reflect.Value() 获取类型；通过kind判断类型
func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Logf("类型: %s, 值： %v \n", reflect.TypeOf(f), reflect.ValueOf(f))
	t.Logf("根据值来判断类型：%s", reflect.ValueOf(f).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	// t.kind() 获取类型，为什么不用reflect.TypeOf(),因为case不能使用float等特殊字符
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Its float")
	case reflect.Int, reflect.Int64:
		fmt.Println("Its int")
	default:
		fmt.Println("Unkown !", t)
	}
}

func TestTypeValue(t *testing.T) {
	var f float64 = 12
	CheckType(f)
}
