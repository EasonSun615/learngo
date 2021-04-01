package main

import (
	"fmt"
)

func arrayTest() {
	// array类型包括array的大小，例如[3]string是一种类型，[5]string是另一种类型
	arr1 := [3]int{1, 2, 3}
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr1) // output:[3]int
	fmt.Printf("%T\n", arr2) // output:[5]int

	// array可以使用==和!=运算符进行比较，所以可以作为map的key
	arr3 := [3]int{1, 2, 3}
	if arr1 == arr3 {
		fmt.Println("true")
	}

	oldArr := [3]int{1, 2, 3}
	newArr := oldArr // 数组是值传递，直接复制一个数组。
	newArr[0] = 5
	fmt.Println(newArr, oldArr)
}

func sliceTest() {
	// slice作为函数参数，会将slice的元数据(metadata describing the slice)复制一份，由于slice的元数据包含指向底层数组的指针，所以看起来像是引用传递
	oldSlice := []int{1, 2, 3}
	fmt.Println(len(oldSlice), cap(oldSlice))
	newSlice := oldSlice // 切片底层基于数组，此时两个切片指向同一个底层数组
	newSlice[0] = 4      // 改变底层数组，所以两个切片都会改变
	fmt.Println(newSlice, oldSlice)
	newSlice = append(newSlice, 5) // 新的切片进行扩容，指向新的底层数组，旧切片仍然指向原来的底层数组
	fmt.Println(newSlice, oldSlice)

	// oldSlice2 := []int{1,2,3}
	// if oldSlice == oldSlice2 {			// slice没有==和!=操作，所以不能作为map的key
	//
	// }
}

func main() {
	// 注意变量和内存的对应关系，
	// 数组类型的变量就对应内存中的一个数组，
	// 切片类型的变量对应的是一个含有三个元素的内存空间，一个是指向对应的底层数组的指针，另外两个是length和capacity(同STL中的vector)
	arrayTest()
	sliceTest()
}

/* 输出：
[3]int
[5]int
true
[5 2 3] [1 2 3]
3 3
[4 2 3] [4 2 3]
[4 2 3 5] [4 2 3]
*/
