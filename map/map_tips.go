package main

import "fmt"

func main() {
	// array 定义
	// 第一种方式，使用大括号，可以同时给出初值
	m1 := map[string]string{}
	fmt.Printf("%v\n", m1)

	// 第二种方式，使用make函数创建map，不能给初值
	m2 := make(map[string]string)
	m2["m2"] = "v2"
	fmt.Printf("%v\n", m2)

	// map的key类型需要定义了==和!=操作, 因为在哈希表的链中要判断是否是已存在的key
	// 所以array可以作为key, slice不行

	// go不同于python, 取一个不存在的key，会返回value类型的空值，不会抛出异常。可以使用ok来判断key是否存在
	a := m1["a"]   // 不会抛出异常
	fmt.Println(a) // a是string的空值--空字符串""

	a, ok := m1["a"]
	if ok {
		fmt.Println("exist key, value=", a)
	} else {
		fmt.Println("not exist key")
	}

	delete(m1, "sun") // 删除元素，注意：删除一个不存在的key，不会抛出异常

}

/*
输出：
map[]
map[m2:v2]

not exist key

*/
