package main

import "fmt"

// func variableZeroValue() {
// 	var a int
// 	var s string
// 	fmt.Printf("%d %q", a, s)
// }

const (
	// 常量组中，后面的如果给出类型和值的话，类型和值会复用前一个常量的类型和表达式。
	Book = iota
	Close
	Phone
)

func main() {
	// variableZeroValue()
	fmt.Println(Book, Close, Phone)
	fmt.Println()
	// iota叫做常量计数器, 主要用来枚举（枚举的时候，常量的值不重要，只要不同就行）
	// 	1.iota只能在常量组中使用
	// 2.不同的const定义块互不干扰
	// 3.没有表达式的常量定义复用上一行的表达式
	// 4. 从第一行开始，iota从0逐行加一
	const (
		a = iota // iota = 0
		b = 10   // iota = 1
		c        // c=10, 复用上一个b的表达式，iota仍然加一，iota=2
		d = iota // iota = 3
	)
	fmt.Print(a, b, c, d)
}
