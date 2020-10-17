package ch4_test

import (
	"fmt"
	"testing"
)

// go语言默认是不支持继承的
type Pet struct {
	Name string
}

func (p *Pet) Speak() {
	fmt.Printf("%s: ....\n", p.Name)
}

func (p *Pet) SpeakTo(someBody string) {
	p.Speak()
	fmt.Printf("%s speak to %s: Hello!\n", p.Name, someBody)
}

type Dog struct {
	p *Pet
}

// 方式一
func TestFuncx(t *testing.T) {
	var p_obj *Pet = &Pet{Name: "Wang"} // 等价于var p_obj Pet = Pert
	var d *Dog = &Dog{p: p_obj}
	d.p.Speak()
}

// 第二种继承方式
type Animal struct {
}

func (a *Animal) Speak() {
	fmt.Println(" ....")
}

func (a *Animal) SpeakTo(someBody string) {
	a.Speak()
	fmt.Printf("speak to %s: Hello!\n", someBody)
}

type Cat struct {
	a *Animal
}

func (c *Cat) Speak() {
	c.a.Speak()
}

func (c *Cat) SpeakTo(someBody string) {
	c.a.SpeakTo(someBody)
}

// 方式二
func TestDemoFunc(t *testing.T) {
	var d = new(Cat)
	d.SpeakTo("Wu")
}
