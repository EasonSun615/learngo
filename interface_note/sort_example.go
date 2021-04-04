package main

import (
	"fmt"
	"sort"
)

type Course struct {
	name  string
	price int
}

// Len() int
// // Less reports whether the element with
// // index i should sort before the element with index j.
// Less(i, j int) bool
// // Swap swaps the elements with indexes i and j.
// Swap(i, j int)

type Courses []Course

func (c Courses) Len() int {
	return len(c)
}

func (c Courses) Less(i, j int) bool {
	return c[i].price < c[j].price
}

func (c Courses) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	var c = Courses{
		Course{"golang", 500},
		Course{"python", 200},
		Course{"music", 999},
	}
	sort.Sort(c) // Sort函数的声明：func Sort(data Interface)， Interface是一个接口（协议）
	for _, course := range c {
		fmt.Printf("%v\n", course)
	}
}
