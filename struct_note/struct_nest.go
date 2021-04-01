package main

import "fmt"

type Teacher struct {
	name string
	age  int
}

// go结构体中不能定义函数，只能定义变量
type Course2 struct {
	name    string
	price   int
	url     string
	Teacher // 匿名嵌套  不用指定变量名
}

func (c Course2) courseInfo() {
	// 嵌套结构体内的不重名属性和普通属性一样访问  重名属性需要指定嵌套的结构体
	fmt.Println("名称：", c.name, "价格：", c.price, "教师信息：", c.Teacher.name, c.age)
}

func main() {
	c := Course2{
		name:  "go",
		price: 99,
		url:   "dd",
		Teacher: Teacher{
			name: "Eason",
			age:  46,
		},
	}
	c.courseInfo() // 名称： go 价格： 99 教师信息： Eason 46

}
