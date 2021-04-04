package main

import (
	"fmt"
	"unsafe"
)

// 接口是一个协议
type Programmer interface {
	Coding()
	Debug() string
}

type Pythoner struct {
}

func (p Pythoner) Coding() {
	fmt.Println("python 开发")
}

func (p Pythoner) Debug() string {
	return "Python debug Done"
}

type Go struct {
}

func (g Go) Coding() {
	fmt.Println("go 开发")
}

func (g Go) Debug() string {
	return "Go Debug Done"
}

func main() {
	// 接口是一种类型，但是它没有数据，不占用内存，是一种抽象类型，不能实例化
	// 接口可以看作编译器为了提供多态的特性而提供的一种占位符
	// 接口个的一个默认的命名规范：接口的名称一般以er结尾
	var pros []Programmer
	pros = append(pros, Pythoner{})
	pros = append(pros, Go{})

	p := Pythoner{}
	fmt.Printf("%T\n", p)         // main.Pythoner
	fmt.Println(unsafe.Sizeof(p)) // 0
	var pro Programmer = Pythoner{}
	fmt.Printf("%T\n", pro) // main.Pythoner
	var pro2 Programmer = Go{}
	fmt.Printf("%T\n", pro2) // main.Go

	// error也是一个接口
	// error的正确使用姿势
	s := "文件不存在"
	var err error = fmt.Errorf("错误：%s", s) // 错误：文件不存在
	fmt.Println(err)

}
