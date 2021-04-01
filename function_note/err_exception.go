package main

import (
	"errors"
	"fmt"
	"time"
)

func div_(a, b int) (int, error) {
	// 写函数时，程序员可以预知到可能会出现除数为零的错误
	if b == 0 {
		return 0, errors.New("除数必须为整数")
	} else {
		return a / b, nil
	}
}

func f1() {
	/*
		子协程抛出的异常不会被父协程捕获到
		一个协程抛出异常但是没有被捕获到的话，整个进程都会挂掉
		所以为了防止进程挂掉，必须在每个子协程中单独捕获异常
	*/
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获到异常")
		}
	}()

	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("捕获到异常2")
			}
		}()
		panic("出错了")
	}()
	time.Sleep(time.Second * 3)
}

func main() {
	/*
		错误：可以预知到的函数会出问题的地方，比如：参数通不过检查，数据库访问不了
		异常是程序员没有考虑到的bug,
	*/
	/*
		 	go 使用panic()抛出异常，在defer后面的函数里使用recover来捕获异常
			坑：子协程抛出的异常不会被父协程捕获到，一个协程抛出异常但是没有被捕获到的话，整个进程都会结束
	*/
	f1()

}
