package main

import "fmt"

func main() {
	/* 在编程中遇到的第一个问题：变量交换
	var a int = 100
	var b int = 200

	var t int
	t = a
	a = b
	b = t
	*/
	// 在Go语言中，程序变量交换，也有语法糖
	var a int = 100
	var b int = 200
	// fmt.Println 可以传递多个参数，用逗号隔开，直接打印
	fmt.Println("a=", a, ",b=", b)
	// 把a,b赋值给b,a  语法糖, 底层本质还是用到了临时变量。简化我们的开发
	b, a = a, b
	fmt.Println("交换后 a=", a, ",b=", b)

	// 复杂的问题都给我们简单化了，我们开发很轻松，编译器帮我们在底层处理。
}
