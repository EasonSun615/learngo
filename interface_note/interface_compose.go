package main

import "fmt"

type Programmer2 interface {
	Coding()
	Debug() string
}

type Designer interface {
	Design() string
}

type Manager interface {
	Programmer2
	Designer
	Manage()
}

type Pythoner2 struct {
	UIDesigner
}

func (p Pythoner2) Coding() {
	fmt.Println("python 开发")
}

func (p Pythoner2) Debug() string {
	return "Python debug Done"
}

type UIDesigner struct {
}

func (u UIDesigner) Design() string {
	fmt.Println("I am a ui designer")
	return "ui designer"
}

func (p Pythoner2) Manage() {
	fmt.Println("我也会管理")
}

func main() {
	// 接口支持组合；结构体组合实现了接口的所有方法也是可以的
	// Manager是一个组合的接口，Python2是一个组合的结构体，只要Python2和其组合的对象实现了所有的接口方法就可以赋值
	var m Manager = Pythoner2{}
	fmt.Printf("%T\n", m) // main.Pythoner2

}
