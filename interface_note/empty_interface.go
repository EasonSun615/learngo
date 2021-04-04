package main

import (
	"fmt"
)

type Course struct {
	name  string
	price int
}

type EmptyInterface interface {
}

func print(i interface{}) {
	fmt.Printf("%v\n", i)
}

func main() {
	// 空接口
	// var ei EmptyInterface
	var i interface{} // 空接口，这是个语法糖：和上面的变量ei完全等价，
	// 空接口可以类似于python中的object, 因为空接口中没有函数，相当于任何结构体都实现了空接口，所以空接口可以接受任何类型的变量
	// 空接口的第一个用途：可以把任何类型都赋值给空接口变量
	// 空接口的第二个用途：参数传递
	i = Course{"chinese", 100}
	fmt.Printf("%T\n", i) // main.Course
	print(i)
	print1(i) // 类型断言
	i = 19
	fmt.Printf("%T\n", i) // int
	print(i)
	print1(i) // 类型断言
	i = []string{"sun", "yi"}
	fmt.Printf("%T\n", i) // []string
	print(i)
	print1(i) // 类型断言

	// 空接口的第三个用途：空接口可以作为map的值
	var teacherInfo = make(map[string]interface{})
	teacherInfo["name"] = "Eason"
	teacherInfo["age"] = 21
	teacherInfo["weight"] = 66.5
	teacherInfo["course"] = []string{"music", "python"}
	fmt.Printf("%v\n", teacherInfo)

}

// 类型断言
func print1(x interface{}) {
	// x.(int) 试图将x转成int型
	v, ok := x.(int)
	if ok {
		fmt.Printf("%d (整数) \n", v)
	}
	// 简介一点的写法
	if s, ok := x.(string); ok {
		fmt.Printf("%s (字符串)\n", s)
	}
	// 更简洁的写法，类型断言的正确姿势：
	switch a := x.(type) {
	case float64:
		fmt.Printf("%f (浮点数）\n", a)
	case []string:
		fmt.Printf("%v (切片) \n", a)
	case Course:
		fmt.Printf("%v (Course结构体）\n", a)
	}

}
