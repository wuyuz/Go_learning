package ch13_test

import (
	"fmt"
	"reflect"
	"testing"
)

// reflect包，定义了两个重要的类型 Type 和 Value 任意接口值在反射中都可以理解为由 reflect.Type 和 reflect.Value
// 两部分组成，并且 reflect 包提供了 reflect.TypeOf和reflect.ValueOf 两个函数来获取任意对象的 Value 和 Type。

// 一、反射的类型对象：reflect.Type？
func TestTypeDome(t *testing.T) {
	var a int
	// 通过 reflect.TypeOf() 取得变量 a 的类型对象 typeOfA，类型为 reflect.Type()。
	typeOfA := reflect.TypeOf(a)
	// 通过 typeOfA 类型对象的成员函数，可以分别获取到 typeOfA 变量的类型名为 int，种类（Kind）为 int。
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	type MyInt int
	var b MyInt
	typeOfB := reflect.TypeOf(b)
	// 结果为：MyInt,int;	表示Name返回实际类型，Kind返回底层类型
	fmt.Println(typeOfB.Name(), typeOfB.Kind())

	typeOfB1 := reflect.TypeOf(&b)
	// 空， ptr
	fmt.Println(typeOfB1.Name(), typeOfB1.Kind())
}

// 二、区分Type和Kind？
// 编程中，使用最多的是类型，但在反射中，当需要区分一个大品种的类型时，就会用到种类（Kind）。例如需要
// 统一判断类型中的指针时，使用种类（Kind）信息就较为方便。
//   1) 反射种类（Kind）的定义：Go语言程序中的类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，
//以及使用 type 关键字定义的类型，这些类型的名称就是其类型本身的名称。例如使用 type A struct{} 定义结构体时，A就是struct{}的类型。
//Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。type A struct{} 定义的结构
//体属于 Struct 种类，*A 属于 Ptr。

func TestSliceDemo(t *testing.T) {
	type myStruct struct{}
	c := myStruct{}
	typeOfC := reflect.TypeOf(c)
	// myStruct struct， 也就是说typeOf得到的额类型对象.Name获取实际类型名，.Kind获取底层类型
	fmt.Println(typeOfC.Name(), typeOfC.Kind())

	typeOfD := reflect.TypeOf(&c)
	// 空， ptr  使用的还是指针
	fmt.Println(typeOfD.Name(), typeOfD.Kind())
}

// 三、使用反射获取结构体成员类型信息
// 任意值通过 reflect.TypeOf() 获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象 reflect.Type 的 NumField()
// 和 Field() 方法获得结构体成员的详细信息。

func TestFunctionDemo(t *testing.T) {
	type cat struct {
		Name string
		// 带有结构体tag的字段
		Type int `json:"type" id:"100"`
	}
	// 创建cat的实例
	ins := cat{Name: "mimi", Type: 1}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 遍历结构体所有成员，  反射类型对象.NumField()获取字段总数
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型, Field()方法返回的结构不再是 reflect.Type 而是 StructField 结构体。
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag， 根据具体的字段.Name和Tag来输出具体的信息
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从tag中取出需要的tag, 使用 StructField 中 Tag 的 Get() 方法，根据 Tag 中的名字进行信息获取。
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}

// 四、结构体Tag
// 通过 reflect.Type 获取结构体成员信息 reflect.StructField 结构中的 Tag 被称为结构体标签（StructTag）。结构体标签是对结构体字段的额外信息标签。
// 1、结构体标签的格式：`key1:"value1" key2:"value2"`（注意没有空格）
// 2、从结构体标签中获取值：
//      2.1、func (tag StructTag) Get(key string) string：根据键获取值，
//      2.2、func (tag StructTag) Lookup(key string) (value string, ok bool)：根据 Tag 中的键，查询值是否存在。

// Type 和 Value 都有一个名为 Kind 的方法，它会返回一个常量，表示底层数据的类型
func TestInterfaceDemo(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float()) // Float 方法用来提取 float64
}

// 可以通过 reflect.Elem() 方法获取这个指针指向的元素类型。这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作
func TestElemDemo(t *testing.T) {
	// 声明一个空结构体
	type cat struct {
	}
	// 创建cat的实例
	ins := &cat{}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 显示反射类型对象的名称和种类
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())
	// 取类型的元素
	typeOfCat = typeOfCat.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
}

func TestProDemo(t *testing.T) {
	// 定义结构体
	type dummy struct {
		a int
		b string
		// 嵌入字段
		float32
		bool
		next *dummy
	}

	// 值包装结构体
	d := reflect.ValueOf(dummy{
		next: &dummy{},
	})
	// 获取字段数量
	fmt.Println("NumField", d.NumField())
	// 获取索引为2的字段(float32字段)
	floatField := d.Field(2)
	// 输出字段类型
	fmt.Println("Field", floatField.Type())
	// 根据名字查找字段
	fmt.Println("FieldByName(\"b\").Type", d.FieldByName("b").Type())
	// 根据索引查找值中, next字段的int字段的值,[]int{4,0} 中的 4 表示，在 dummy 结构中索引值为 4 的成员，也就是 next。next 的类型为 dummy，也是一个结构体，
	// 因此使用 []int{4,0} 中的 0 继续在 next 值的基础上索引，结构为 dummy 中索引值为 0 的 a 字段，类型为 int。
	fmt.Println("FieldByIndex([]int{4, 0}).Type()", d.FieldByIndex([]int{4, 0}).Type())
}

// 反射值对象（reflect.Value）提供一系列方法进行零值和空判定，
