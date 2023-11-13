package main

import "fmt"

func main() {
	//
	var num int = 15

	// if 不满足则跳过该if里面的代码，程序不会执行
	if num > 100 {
		fmt.Println("num > 100")
	}
	// if的代码必须要条件满足才可以执行
	if num > 10 {
		fmt.Println("num > 10")
	}
	fmt.Println("main end")
}
