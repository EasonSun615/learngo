package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 指针的空值是nil
	// go中的指针相对于C/C++中的指针，做了大量的限制，只能用来读写值，不能做指针的转换、偏移、运算，这样的安全性高

	// make new nil
	// 每个变量声明的时候没有初始化的话，都会有一个默认值
	// int byte float bool string 等这些类型的默认值是对应的零值（0， ""， false)
	// 指针，slice, map的默认值是nil, 这些类型都是指向了别的内存地址
	//
	var amap map[string]string
	var aslice []string
	var aptr *string
	fmt.Println(unsafe.Sizeof(amap), unsafe.Sizeof(aslice), unsafe.Sizeof(aptr))
	var aarray [3]int32
	fmt.Println(unsafe.Sizeof(aarray))

}
