package main

import "fmt"

// 逻辑运算符，结果也是布尔值。 多个条件判断是否成立作用
func main() {
	var a bool = true
	var b bool = false
	var c bool = true
	var d bool = false

	// 双元运算符  &&   ||

	// 与 都为真结果才为真，如果有假则结果为假
	// 账户名 和 密码 都要正确才可以登录
	fmt.Println(a && b)
	fmt.Println(a && c)
	fmt.Println(a && d)
	// 短路判断, 在这里 4<5 是没有运算的，也不要运算，因为在 3<2 结果已经出现了
	fmt.Println(1 == 1 && 3 < 2 && 4 < 5)
	fmt.Println("=============================")
	// 或 || 只要满足一个条件，整体结果就位 true  or
	// 如果所有条件都不满足，结果为false。
	// 看文章 （登录、缓存、能够访问......）
	fmt.Println(a || b) // true  false
	fmt.Println(a || c) // true  true
	fmt.Println(a || d) // true  false
	fmt.Println(b || d) // false  false

	// 单元运算符  一个条件就够，就是将当前结果的bool置反
	// 如果真，则结果是假，如果假，则结果是真，
	// 取反
	fmt.Println(!(b || d))
	fmt.Println(!true)
	fmt.Println(!!false)
	//
	// 结果是false
	if !false { //true
		// 执行...
	}

}
