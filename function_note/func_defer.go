package main

import "fmt"

func main() {
	/*
		defer 用来解决资源的释放问题
		defer语句是go语言的一种注册延迟调用的机制，
		defer之后只能是函数调用
		defer语句会被压入一个栈中，所以后面的defer语句先执行
	*/
	fmt.Println("test1")
	defer fmt.Println("defer test1")
	fmt.Println("test2")
	defer fmt.Println("defer test2")
	fmt.Println("test3")
	defer fmt.Println("defer test3")
}

/*
输出：
test1
test2
test3
defer test3
defer test2
defer test1

*/
