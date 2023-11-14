package main

import "fmt"

// 函数作用域
// 局部变量
// 1、函数内部定义的变量，只能在函数内部调用
// 全部变量(全局变量)
// 1、函数外部定义的变量,默认我们定义在上面，方便文件统一查看和管理全局变量
var num int = 1

func main() {
	// 定义一个只在 if 中生效的变量 if 临时变量(a,b := 1,2);条件判断{}
	// 小作用域可以用到大作用域中的变量，反之则不行。
	// 对于很多一次性的变量，都可以这么写
	// 优先考虑使用局部变量。
	if a := 11; a <= 10 {
		fmt.Println(num)
		fmt.Println(a)
		// 就近原则
		if a := 2; a <= 10 {
			fmt.Println(a)
		}
	}
	fmt.Println(num)
}

func f1() {
	fmt.Println(num)
}
func f2() {

}
