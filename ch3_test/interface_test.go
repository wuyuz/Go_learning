package ch3_test

import (
	"fmt"
	"testing"
)

// 定义一个接口，可以理解为所有的结构体都是继承与接口的，当继承的结构体都实现了当前接口的方法，那么就继承了这个接口
// 也就可以通过这个接口去动态/多态的调用接口的方法了，换句话说，如果接口体没有实现接口类，必然会执行到父类接口的函数签名
// 也就是无法实现接口，类似与python中的父类抛错
type Programmer interface {
	SayHello() string
	Hi()
}

// go实例
type GoProgrammer struct {
	Id   string
	Name string
}

func (g *GoProgrammer) SayHello() string {
	fmt.Println("GO Programmer 实现")
	return "Go对象返回"
}

func (g *GoProgrammer) Hi() {
	fmt.Printf("Say: %s Hi\n", g.Name)
}

// py对象
type PyProgrammer struct {
	Name string
}

func (g *PyProgrammer) SayHello() string {
	fmt.Println("Python Programmer 实现")
	return "Python对象返回"
}

func (g *PyProgrammer) Hi() {
	fmt.Printf("Say: %s Hi\n", g.Name)
}

func TestInterface(t *testing.T) {
	var p Programmer      // 声明p为Programmer接口对象
	p = new(GoProgrammer) // 创建一个Goprogrammer结构体的实例指针指向p ,如果为满足，则会抛错
	ret := p.SayHello()   // 由于GoProgrammer实现了签名函数，所以可以调用
	fmt.Println(ret)

	// 鸭子类型：多态，两个结构体都有同样的方法，功能很像就一样
	p = new(PyProgrammer)
	ret = p.SayHello()
	fmt.Println(ret)
}

func TestInterfaceFunc(t *testing.T) {
	// 接口变量， 接口对象 = 结构体类型 + 实例数据，也就是说拥有结构体的方法和数据
	var prog Programmer = &GoProgrammer{Id: "1", Name: "wang"}
	var pyprog Programmer = &PyProgrammer{Name: "li"}
	prog.Hi()
	pyprog.Hi()
}
