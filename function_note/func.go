package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
	} else {
		result = a / b
	}
	return
}

func sum(numbers ...int) (sum int) {
	// 省略号用法二：在形参中，表明接受可变数量的参数
	for _, num := range numbers {
		sum += num
	}
	return
}

func main() {
	res, err := div(6, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	fmt.Println(sum(1, 2, 3))
	s := []int{1, 2, 3, 4}
	fmt.Println(sum(s...)) // 省略号用法二：给实参slice 解容器

	arr := [...]int{1, 2, 3} // 省略号用法三：初始化数组时，可以用省略号来代替数组大小
	fmt.Printf("%T\n", arr)

	// go中函数是一等公民，一等公民就是说函数和普通的变量一样，可以作为函数参数，函数返回值，可以进行赋值
	first_class()

}

func first_class() {
	/*
		函数作为第一公民的特性：
			函数名是一个变量，类型是函数声明。普通类型的变量可以做的操作，函数名也可以
			函数内部可以定义函数，函数内部的函数必须是匿名函数，函数内部的函数必须赋值给一个变量或者在定义的时候直接调用
	*/
	var va, vb int = 111, 222
	// 1、内部匿名函数赋值给一个函数型变量
	innerAdd := func(a, b int) int {
		return a + b
	}
	fmt.Printf("%T\n", innerAdd)  // func(int, int) int
	fmt.Println(innerAdd(va, vb)) // 333
	// 2、内部匿名函数在定义时直接调用
	vc := func(a, b int) int {
		return a + b
	}(va, vb)
	fmt.Printf("%T\n", vc) // int
	fmt.Println(vc)        // 333
}

/*
输出：
除数不能为0
6
10
[3]int
func(int, int) int
333
int
333
*/
