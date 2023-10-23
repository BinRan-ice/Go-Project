package main

import "fmt"

// byte
func main() {
	// 别名uint8是byte，十分常用
	// byte = uint8 （0-255之间的整数，我们通常使用byte来定义） uint8
	var num1 byte = 255
	fmt.Println(num1)
	fmt.Printf("%T\n", num1)

	// 不经常用 rune int32
	var num2 rune = 1000000000
	fmt.Println(num2)
	fmt.Printf("%T\n", num2)

	// int、系统大小来的，32位  32、64默认是64
	// 软件跑在32位系统上是不兼容的
	var num3 int = 100000
	fmt.Println(num3)
	fmt.Printf("%T\n", num3)
}
