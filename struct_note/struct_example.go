package main

import "fmt"

// go结构体中不能定义函数，只能定义变量
type Course struct {
	name  string
	price int
	url   string
}

func (c Course) printCourseName() {
	fmt.Println("名称：", c.name)
}

func (c Course) setPriceWrong(p int) {
	c.price = p
}

func (c *Course) setPrice(p int) {
	c.price = p // 语法糖 相当于语句： （*c).price = p
}

func structExample() {
	// go语言没有class概念
	// go语言中结构体名和变量名的首字母大小写决定着结构体或者变量的可见范围
	// 结构体用大括号实例化，
	// 结构体实例化方式1：k-v模式
	var c1 Course = Course{
		name:  "django",
		price: 100,
		url:   "www.imooc.com",
	}
	fmt.Println(c1)
	// 结构体实例化方式2：顺序模式
	c2 := Course{"python", 100, "www.dd"}
	fmt.Println(c2)
	c3 := &Course{"golang", 99, "www.go.com"}
	// 语法糖就是编译器为了方便写代码而设置的一些语法
	fmt.Println(c3.name, c3.price, c3.url) // 这里是go语言的一个语法糖，结构体指针可以直接像结构体一样访问元素，这里编译器会将c3.name转成（*c3).name
}

func structFunctionExample() {
	/*
			总结：结构体的方法必须和结构体在同一个包内
		 		 结构体的方法尽量绑定结构体的指针，指针传递才可以修改对象，并且可以防止产生复制结构体的开销
	*/
	c := Course{name: "golang", price: 88, url: "www.golang.cn"}
	c.printCourseName() // 这里是一个语法糖  相当于语句：Course.printCourseName(c)
	c.setPriceWrong(99) // 语法糖 相当于语句：Course.setPriceWrong(c, 99)
	Course.setPriceWrong(c, 99)
	fmt.Println(c.price) // 结构体是值传递，所以不会改变结构体
	c.setPrice(66)       // 语法糖 相当于语句：(&c).setPrice(66)  ->  (*Course).setPrice(&c, 66)
	fmt.Println(c.price) // 66
	(*Course).setPrice(&c, 77)
	fmt.Println(c.price) // 77
}

func main() {
	structExample()
	structFunctionExample()
	// go 通过匿名嵌套来实现继承的效果
}
