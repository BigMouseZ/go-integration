package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

// 在reflect包中，最重要的两个类型就是Type和Value，分别从类型、值的角度来描述一个Go对象。
type testStruct struct {
	NameTest string
	AgeTest  int
	FileTest *os.File
	PtrTest  *string
}

func TestType(t *testing.T) {
	argTest := "test"
	v := reflect.TypeOf(argTest)

	fmt.Println(v.Kind()) // string

	argTest1 := &testStruct{}
	v1 := reflect.TypeOf(argTest1)

	fmt.Println(v1.Kind()) //  ptr

	argTest1 = &testStruct{}
	v1 = reflect.TypeOf(*argTest1)

	fmt.Println(v1.Kind()) // struct

	argTest1 = &testStruct{}
	v1 = reflect.TypeOf(argTest1).Elem()

	fmt.Println(v1.Kind()) // struct
}

// 获取结构体中所有元素的属性
func TestStructType(t *testing.T) {
	argTest1 := &testStruct{}
	getStructArgProperty(argTest1)
}
func getStructArgProperty(t interface{}) {
	var v reflect.Type
	if reflect.TypeOf(t).Kind() == reflect.Ptr {
		if reflect.TypeOf(t).Elem().Kind() != reflect.Struct {
			fmt.Println("不是结构体")
			return
		}
		v = reflect.TypeOf(t).Elem()
	} else {
		if reflect.TypeOf(t).Kind() != reflect.Struct {
			fmt.Println("不是结构体")
			return
		}
		v = reflect.TypeOf(t)
	}
	run(v)
}
func run(v reflect.Type) {
	for i := 0; i < v.NumField(); i++ {
		argType := v.Field(i)
		if argType.Type.Kind() == reflect.Ptr {
			fmt.Println(argType.Name, argType.Type.Elem().Kind())
		} else {
			if argType.Type.Kind() == reflect.Struct {
				fmt.Println("   =====>", argType.Name)
				run(argType.Type)
			} else {
				fmt.Println(argType.Name, argType.Type.Kind())
			}
		}
	}
}

// 获取结构体中所有元素的属性
func TestValue(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}
	v := reflect.ValueOf(arr)
	fmt.Println(v) // [1,2,3,4]

	v1 := reflect.ValueOf(&arr)
	fmt.Println(v1) // &[1,2,3,4]

	// fmt.Println(v.Elem().CanSet())  // panic
	fmt.Println(v1.Elem().CanSet()) // true

	v1.Elem().Index(0).SetInt(10)
	fmt.Println(arr) // 10,2,3,4
}

type student struct {
	numb  int
	name  string
	Age   int
	class *class
}
type class struct {
	classNumber int
	className   string
}

func TestStructValueOf(t *testing.T) {
	s := student{numb: 1, name: "john", Age: 18, class: &class{classNumber: 1}}
	v := reflect.ValueOf(&s)
	getStructArgPropertyValue(v)
}

func getStructArgPropertyValue(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ { // NumField()会判断Kind()是否为struct 不是的话会panic
		argType := v.Field(i)
		if argType.Kind() == reflect.Ptr {
			if argType.Elem().Kind() == reflect.Struct {
				fmt.Println("========>")
				getStructArgProperty(argType.Elem())
			} else {
				fmt.Println(argType.Elem().Kind(), "     : ", argType, "   ", argType.Elem().CanSet())
			}
		} else {
			if argType.Kind() == reflect.Struct {
				getStructArgProperty(argType)
			} else {
				if argType.CanSet() == true && argType.Kind() == reflect.Int {
					argType.SetInt(10)
				}
				fmt.Println(argType.Kind(), "     : ", argType, "   ", argType.CanSet())
			}
		}
	}
}

// 若要获取类型的方法，使用TypeOf(),ValueOf()2中类型都可以获取。
//
// 不同的是TypeOf()返回方法的基本属性，但并自己没有现实调用方法，而是通过调用ValueOf的Call(),而ValueOf则没有返回方法的名字等基本属性

type myType int

func (my *myType) Hi() {
	fmt.Println("my value ", *my)
}
func (my *myType) Set(x int) {
	*my = myType(x)
}
func (my myType) Get() int {
	fmt.Println("my value ", my)
	return int(my)
}

var s myType = 1

func TestGetSetStruct(t *testing.T) {
	v := reflect.ValueOf(&s)
	v1 := reflect.TypeOf(s)

	fmt.Println(" v  ", v.NumMethod())   // 3
	fmt.Println(" v1  ", v1.NumMethod()) // 1  传入的如果是值类型，则只返回值类型方法

	for i := 0; i < v1.NumMethod(); i++ {
		fmt.Println(v1.Method(i)) // 方法名等结果，根据首字母排序
	}

	for i := 0; i < v.NumMethod(); i++ {
		fmt.Println(v.Method(i)) // reflect方法对象。
	}

	var para []reflect.Value
	para = append(para, reflect.ValueOf(11))
	fmt.Println(v.Method(2).Call(para)) // 调用Set方法

	para = append(para, reflect.ValueOf(&s))
	fmt.Println(v1.Method(0).Func.Call(para[1:])) // 调用Get方法
}
