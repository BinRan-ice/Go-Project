package main

import "fmt"

func main() {
	// 算术运算符 + - * / % ++ --
	var a int = 7
	var b int = 3
	var c int // 结果
	// 双元运算符，由两个变量参与运算，或者多个
	c = a + b
	fmt.Println(c)
	c = a - b
	fmt.Println(c)
	c = a * b
	fmt.Println(c)

	c = a / b // 除 7/3 = 2...1
	fmt.Println(c)
	c = a % b // 取模 余数  1
	fmt.Println(c)

	// 单元运算符，用自己就可以操作得出结果
	//  自增++  自减--
	fmt.Println(a)
	a++ // a = a + 1 自增，在自己的基础上 + 1
	fmt.Println(a)

	a = 7
	a-- // a = a-1。  遍历输出一些东西（10000个数字）
	fmt.Println(a)

}
