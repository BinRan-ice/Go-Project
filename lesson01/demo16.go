package main

import "fmt"

func main() {

	var str string
	str = "Hello,World"
	fmt.Printf("%T,%s\n", str, str)
	// H E L L O , W O R L D
	// Go语言中，所有的字符串都是由单个 字符 连接起来的。
	// 单引号是字符、双引号才是 string 类型
	// 字符本质是整型
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T,%d\n", v1, v1)
	fmt.Printf("%T,%s\n", v2, v2)

	v3 := '魑'
	fmt.Printf("%T,%d\n", v3, v3)
}
