package main

import (
	"fmt"
	"unsafe"
)

func funcName(a int) bool {
	if a > 10 {
		return true
	} else {
		return false
	}
}

func main() {
	/*
		Go对变量的类型卡的很死，即使是内存结构完全相同的类型，变量之间也不可以互相赋值
	*/
	type tt int
	var va tt = 10
	// var vb int = va		报错：cannot use va (type tt) as type int in assignment
	fmt.Println(va)
	fmt.Println(unsafe.Sizeof(va))
	var vb int
	fmt.Println(unsafe.Sizeof(vb))
	// 数组的变量类型包含元素类型和元素个数
	arr1 := [2]int{1, 2}
	arr2 := [3]int{1, 2, 3}
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)
	// 函数的变量类型就是函数的声明
	fmt.Printf("%T\n", funcName) // func(int) bool
	a := funcName
	fmt.Printf("%T\n", a) // func(int) bool

}

/*
输出：
10
8
8
[2]int
[3]int
func(int) bool
func(int) bool
*/
