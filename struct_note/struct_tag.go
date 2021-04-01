package main

import (
	"fmt"
	"reflect"
)

type Player struct {
	name string `database:"name, max_length=17"`
	age  int    `database:"age min=18"`
}

func main() {
	p := Player{
		name: "Eason",
		age:  22,
	}
	t := reflect.TypeOf(p)
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind", t.Kind())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("database")
		fmt.Printf("%d. %v (%v), tag:'%v'\n", i, field.Name, field.Type, tag)
	}

}

/*
输出：
Type: Player
Kind struct
0. name (string), tag:'name, max_length=17'
1. age (int), tag:'age min=18'
*/
