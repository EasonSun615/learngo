package main

import "fmt"

func filterSlice(s []int, f func(int) bool) (out []int) {
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return
}

func filterFunc(a int) bool {
	if a >= 70 {
		return true
	} else {
		return false
	}
}

func main() {
	// 函数作为一等公民的特性：作为另外一个函数的参数
	score := []int{90, 70, 40, 50, 60}
	// 方式一：使用匿名函数
	access := filterSlice(score, func(i int) bool {
		if i >= 60 {
			return true
		} else {
			return false
		}
	})
	fmt.Println(access)
	// 方式二：直接使用函数名
	access2 := filterSlice(score, filterFunc)
	fmt.Println(access2)
}
