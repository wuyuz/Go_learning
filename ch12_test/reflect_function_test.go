package ch12_test

import (
	"fmt"
	"reflect"
	"testing"
)

// 学习网址： https://www.cnblogs.com/qcrao-2018/p/10822655.html

// 一、通过反射实现灵活的功能？
// 1、 按名字访问结构的成员： reflect.ValueOf(*e).FieldByName("Name")
// 2、 按名字访问结构的方法： reflect.value(e).MethodByName("updateAge").Call([]reflect.Value{reflect.ValueOf(1)})

// 二、 反射的定义？
// 反射：在计算机科学中，反射是指计算机程序在运行时（Run time）可以访问、检测和修改它本身状态或行为的一种能力。用比喻来说，反射就是程序在运行
//的时候能够“观察”并且修改自己的行为； 《go 圣经》中定义：Go 语言提供了一种机制在运行时更新变量和检查它们的值、调用它们的方法，但是在编译时并
//不知道这些变量的具体类型，这称为反射机制。

// 三、不用反射就不能在运行时访问、检测和修改它本身的状态和行为吗？
//   答：首先理解什么叫访问、检测和修改它本身状态或行为，它的本质是什么？实际上，它的本质是程序在运行期探知对象的类型信息和内存结构，不用反射能行吗？
//  可以的！使用汇编语言，直接和内层打交道，什么信息不能获取？但是，当编程迁移到高级语言上来之后，就不行了！就只能通过反射来达到此项技能。

// 四、为什么要用反射？（1、参数未知；2、动态调用函数）
// 1、有时你需要编写一个函数，但是并不知道传给你的参数类型是什么，可能是没约定好；也可能是传入的类型很多，这些类型并不能统一表示。这时反射就会用的上了。
// 2、有时候需要根据某些条件决定调用哪个函数，比如根据用户的输入来决定。这时就需要对函数和函数的参数进行反射，在运行期间动态地执行函数。

// 五、反射的缺点？
// 1、难以阅读，水平要求高
// 2、go语言是是一门静态语言，在编译的时候会发现错误，但是对于反射来说就无能为力，因为不知道具体的情况下的结果，所以在一段时候后可能会发现错误
// 3、性能损耗大，比正常代码慢一至两个等级，注意权衡

// 六、 反射的实现原理？
// 基于interface，它是 Go语言实现抽象的一个非常强大的工具。当向接口变量赋予一个实体类型的时候，接口会存储实体的类型信息，反射就是通过接口的类型信息实现的，
// 反射建立在类型的基础上，reflect 包里定义了各种类型，实现了反射的各种函数，通过它们可以在运行时检测类型的信息、改变类型的值。

// 七、types和interface？
// Go 语言中，每个变量都有一个静态类型，在编译阶段就确定了的，比如 int, float64, []int 等等。注意，这个类型是声明时候的类型，不是底层数据类型。
type myInt int

func TestType(t *testing.T) {
	var i int
	var j myInt
	// fmt.Println(i == j)  // 报错，不能强等,尽管底层类型相同，但我们知道，他们是不同的静态类型，除非进行类型转换，否则，i和j不能同时出现在等号两侧。j的静态类型就是 MyInt。
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(j))   // int	xxx.myInt
	fmt.Println(reflect.ValueOf(i), reflect.ValueOf(j)) // 默认初始值0
}

// 反射相关函数练习

type Child struct {
	Name     string
	Grade    int
	Handsome bool
}

type Adult struct {
	ID         string `qson:"Name"`
	Occupation string
	Handsome   bool
}

// 如果输入参数 i 是 Slice，元素是结构体，有一个字段名为 `Handsome`，
// 并且有一个字段的 tag 或者字段名是 `Name` ，
// 如果该 `Name` 字段的值是 `qcrao`，
// 就把结构体中名为 `Handsome` 的字段值设置为 true。
func handsome(i interface{}) {
	// 获取 i 的反射变量 Value
	v := reflect.ValueOf(i)

	// 确定 v 是一个 Slice
	if v.Kind() != reflect.Slice {
		return
	}

	// 确定 v 是的元素为结构体
	if e := v.Type().Elem(); e.Kind() != reflect.Struct {
		return
	}

	// 确定结构体的字段名含有 "ID" 或者 json tag 标签为 `name`
	// 确定结构体的字段名 "Handsome"
	st := v.Type().Elem()

	// 寻找字段名为 Name 或者 tag 的值为 Name 的字段
	foundName := false
	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		tag := f.Tag.Get("qson")

		if (tag == "Name" || f.Name == "Name") && f.Type.Kind() == reflect.String {
			foundName = true
			break
		}
	}

	if !foundName {
		return
	}

	if niceField, foundHandsome := st.FieldByName("Handsome"); foundHandsome == false || niceField.Type.Kind() != reflect.Bool {
		return
	}

	// 设置名字为 "qcrao" 的对象的 "Handsome" 字段为 true
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		handsome := e.FieldByName("Handsome")

		// 寻找字段名为 Name 或者 tag 的值为 Name 的字段
		var name reflect.Value
		for j := 0; j < st.NumField(); j++ {
			f := st.Field(j)
			tag := f.Tag.Get("qson")

			if tag == "Name" || f.Name == "Name" {
				name = v.Index(i).Field(j)
			}
		}

		if name.String() == "qcrao" {
			handsome.SetBool(true)
		}
	}
}

func TestFunctionOfReflect(t *testing.T) {
	children := []Child{
		{Name: "Ava", Grade: 3, Handsome: true},
		{Name: "qcrao", Grade: 6, Handsome: false},
	}

	adults := []Adult{
		{ID: "Steve", Occupation: "Clerk", Handsome: true},
		{ID: "qcrao", Occupation: "Go Programmer", Handsome: false},
	}

	fmt.Printf("adults before handsome: %v\n", adults)
	handsome(adults)
	fmt.Printf("adults after handsome: %v\n", adults)

	fmt.Println("-------------")

	fmt.Printf("children before handsome: %v\n", children)
	handsome(children)
	fmt.Printf("children after handsome: %v\n", children)
}
