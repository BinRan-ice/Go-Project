package main

import "fmt"

// 常量和变量放置的内存地址不同  （栈、堆、常量池）
// 程序正常执行,压栈

// 常量
func main() {
	// 规定：建议这么去做
	// 我们通过定义常量的时候，建议大家使用大写字母来定义。区分与普通变量
	// 一旦定义之后是不会发生变化的。
	// 定义格式  const 常量名[type] = value
	const URL string = "www.kuangstudy.com"

	// 隐式定义 常量的自动推导是可以省略一些基础类型，
	const URL2 = "www.baidu.com"

	// 可以同时定义多个常量
	const URL3, URL4 string = "www.kuangstudy.com", "www.baidu.com"

	//
	fmt.Println(URL)
	fmt.Println(URL2)
	fmt.Println(URL3)
	fmt.Println(URL4)

	// 在我们真实的世界也是有很多不会发生变化量，那在程序中对应的就是常量
	const PI = 3.14
	// 固定的东西，都建议统一定义成常量。
	const LENGTH int = 8000
	const WIDTH int = 8000

	// 常量是无法被修改的。
	//LENGTH = 13

	fmt.Println(LENGTH)

}
