package main

import "fmt"

func main() {
	// 调用函数 ： 函数名 +  () 调用
	printHelloWorld()
	print2()
	print3()
}

// func 定义函数
/**
func 函数名(参数1，参数2,....) {  计算器
	// 封装的思维，代码多了我们将一些功能代码提出来，方便复用，让代码更清晰。
	// 函数的逻辑，本质就是一段代码，我们未来很多业务都会不断的抽离。
}
*/
func printHelloWorld() {
	fmt.Println("Hello,world")
}
func print2() {
	fmt.Println("Hello,world222")
}
func print3() {
	fmt.Println("Hello,world333")
}
