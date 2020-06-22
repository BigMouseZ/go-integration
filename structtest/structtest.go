package main

import "fmt"

// 结构体  for range 指针特性测试

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
		//根本原因在于for-range会使用同一块内存去接收循环中的值。
		/*
			大王八 => 大王八
			小王子 => 大王八
			娜扎 => 大王八
		*/
	}
	//改进
	for _, stu := range stus {
		pointAdress := &stu
		aaa := stu
		fmt.Printf("n.a %p\n", &pointAdress)
		fmt.Printf("n.a %p\n", &aaa)
		m[stu.name] = &aaa
		/*
			大王八 => 大王八
			小王子 => 小王子
			娜扎 => 娜扎
		*/
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}