package ch4_test

import (
	"fmt"
	"testing"
)

type GoProgrammer struct {
}

func (g *GoProgrammer) Say() {
	fmt.Println("fmt.Println(\" Hello \"))")
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) Say() {
	fmt.Println("System.out.Println(\" Hello \")")
}

type Programmer interface {
	Say()
}

// 多态的实现函数不能是指针传递
func PrintProgrammer(p Programmer) {
	p.Say()
}

// 多态方法的实现，不同的实现对象的结果不一样
func TestFuncs(t *testing.T) {
	var g = new(GoProgrammer) // 这里不能使用 var g = GoProgrammer{}, 因为指针对象p只接受指针类型
	var j = new(JavaProgrammer)

	// pointer := &GoProgrammer{}
	var p1 = g
	var p2 = j
	// var p3 = pointer
	p1.Say()
	p2.Say()
	// p3.Say()
}
