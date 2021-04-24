package main

import "fmt"

const (
	// 常量组中，后面的如果没有给出类型和值的话，类型和值会复用前一个常量的类型和表达式。
	BOOK = iota
	CLOSE
	PHONE
)

// 枚举类型的常量的编码规范：
// 1.常量均适用大写字母，使用下划线分词
// 2.如果是枚举类型的常量，需要先创建相应类型。如下所示

type Scheme string

const (
	HTTP  Scheme = "http"
	HTTPS Scheme = "https"
)

func main() {
	fmt.Println(BOOK, CLOSE, PHONE)
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

/*
输出：
0 1 2
0 10 10 3
*/
