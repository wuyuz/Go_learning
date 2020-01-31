package ch3_test

import (
	"fmt"
	"testing"
	"unsafe"
)

// 类似于面向对象
type Employee struct {
	id   string
	name string
	age  int
}

func TestInit(t *testing.T) {
	e := Employee{"0", "wang", 20}
	e1 := Employee{name: "wu", age: 22}

	// 创建一个指向实例的指针，相当于 e3 = &Employee{}，可以对其属性进行赋值；
	// 需要注意的是，GO语言中通过结构体的对象指针和实例的使用方法一致，直接.属性就可以了
	e3 := new(Employee)
	e3.name = "liu"
	e3.age = 23
	t.Log(e)
	fmt.Println(e)
	fmt.Println(e1)
	fmt.Println(e3)
	t.Logf("e1 objdect is %T", e1)
	t.Logf("e3 object is %T", e3)   // 表死指针
	t.Logf("e1 objdect is %T", &e1) // 和上一个一致
}

// 不推荐，在实例方法被调用时，会发生值拷贝
func (e Employee) Strings() string {
	fmt.Println("Address is: ", unsafe.Pointer(&e.name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.id, e.name, e.age)
}

// 推荐，在实例方法被调用时，会不发生值拷贝
func (e *Employee) String() string {
	fmt.Println("Address is: ", unsafe.Pointer(&e.name)) //未发生改变
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.id, e.name, e.age)
}

func TestFunc(t *testing.T) {
	e4 := Employee{"3", "wanglixing", 24}
	fmt.Println("Old Address is: ", unsafe.Pointer(&e4.name))
	fmt.Println(e4.Strings())
	fmt.Println(e4.String()) // 按理说这是函数是给其指针的，但是这里实例也可以调用了
}
