package main

import "fmt"

func main() {
	// type的用途：
	// 1、给一个类型定义别名，增强代码的可读性。例如byte类型是uint8的别名
	// 为什么给uint8定义一个byte的别名，就是为了强调现在处理的对象是字节类型，作用是增强了代码的可读性，编译器在编译时仍然会将byte替换成uint8
	type myByte = byte
	var a myByte
	fmt.Printf("%T\n", a) // uint8

	// 2.基于一个已有的类型定义一个新的类型
	type myInt int
	var b myInt
	fmt.Printf("%T\n", b) // main.myInt

	// 3、定义结构体
	type Course struct {
		name  string
		price int
	}

	// 4、定义接口
	type Callable interface {
	}
}
