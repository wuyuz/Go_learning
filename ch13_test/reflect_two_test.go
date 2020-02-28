package ch13_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsVailDemo(t *testing.T) {
	// *int的空指针
	var a *int
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())
	// nil值,将变量 a 包装为 reflect.Value 并且判断是否为空，此时变量 a 为空指针，因此返回 true。
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())
	// *int类型的空指针,对 nil 进行 IsValid() 判定（有效性判定），返回 false。
	// (*int)(nil) 的含义是将 nil 转换为 *int，也就是*int 类型的空指针。此行将 nil 转换为 *int 类型，并取指针指向元素。
	// 由于 nil 不指向任何元素，*int 类型的 nil 也不能指向任何元素，值不是有效的。因此这个反射值使用 Isvalid() 判断时返回 false。
	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())
	// 实例化一个结构体
	s := struct{}{}
	// 尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid())
	// 尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(s).MethodByName("").IsValid())
	// 实例化一个map
	m := map[int]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
}

func TestAddrDemo(t *testing.T) {
	// 对于 reflect.Values 也有类似的区别。有一些 reflect.Values 是可取地址的；其它一些则不可以。必须使用Elem()
	x := 2                   // value type variable?
	a := reflect.ValueOf(2)  // 2 int no
	b := reflect.ValueOf(x)  // 2 int no
	c := reflect.ValueOf(&x) // &x *int no
	d := c.Elem()            // 2 int yes
	// 我们可以通过调用 reflect.ValueOf(&x).Elem()，来获取任意变量x对应的可取地址的 Value。

	fmt.Println(a.CanAddr()) // "false"
	fmt.Println(b.CanAddr()) // "false"
	fmt.Println(c.CanAddr()) // "false"
	fmt.Println(d.CanAddr()) // "true"
	// Addr() Value获取地址，CanAddr() bool 用于返回
}

// 值可被修改的条件：1、可寻址；2、可被导出
func TestSetDemo(t *testing.T) {
	// 声明整型变量a并赋初值
	var a int = 1024
	// 获取变量a的反射值对象(a的地址)
	valueOfA := reflect.ValueOf(&a)
	fmt.Println("Elem before: ", valueOfA) // 0xc000014130
	// 取出a地址的元素(a的值)
	valueOfA = valueOfA.Elem()
	fmt.Println("Elem after: ", valueOfA) //  1024
	// 修改a的值为1
	valueOfA.SetInt(1)
	// 打印a的值
	fmt.Println("finally val: ", valueOfA.Int()) // 1
}

// 使用反射调用函数：如果反射值对象（reflect.Value）中值的类型为函数时，可以通过 reflect.Value 调用该函数。
//使用反射调用函数时，需要将参数使用反射值对象的切片 []reflect.Value 构造后传入 Call() 方法中，调用完成时，
//函数的返回值通过 []reflect.Value 返回。

// 具体使用步骤：下面的代码声明一个加法函数，传入两个整型值，返回两个整型值的和。将函数保存到反射值对象（reflect.Value）中，然后
//将两个整型值构造为反射值对象的切片（[]reflect.Value），使用 Call() 方法进行调用。

// 普通函数
func add(a, b int) int {
	return a + b
}

func TestReflectFunction(t *testing.T) {
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)
	// 构造函数参数, 传入两个整型值
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	// 反射调用函数
	retList := funcValue.Call(paramList)
	// 获取第一个返回值, 取整数值
	fmt.Println(retList[0].Int())
}

type Cat struct {
	Name  string
	Color string
	Age   int
}

func (p *Cat) Eat() int {
	fmt.Printf("Cat: %s , Its Eatting!\n", p.Name)
	return 6
}

// 唱歌的方法,具有两个入参
func (cat *Cat) Sing(sing string, i int) int {
	fmt.Println(sing)
	fmt.Println(i, "Its singing")
	return 2
}

// 结构体中调用方法
func ReflectMethod(o interface{}) {
	v := reflect.ValueOf(o)
	//无参函数调用
	m1 := v.MethodByName("Eat")
	//调用无参方法
	call := m1.Call([]reflect.Value{})
	for key, value := range call {
		fmt.Println(key, value)
	}
	//有参函数调用
	m2 := v.MethodByName("Sing")
	values := m2.Call([]reflect.Value{reflect.ValueOf("iris"), reflect.ValueOf(3)})
	for _, value := range values {
		fmt.Println(value)
	}
}

func TestReflectDemo(t *testing.T) {
	// 需要使用指针否则会抛错
	ReflectMethod(&Cat{Name: "petty", Color: "red", Age: 12})
}
