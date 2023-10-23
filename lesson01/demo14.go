package main

import "fmt"

func main() {
	// 浮点数 ； 符号位  + 指数位 + 尾数位（存储过程中,可能会丢失，造成精度损失）。
	// float64 的空间 > float32
	var num1 float32 = 123.0000901
	var num2 float64 = 123.0000901
	fmt.Println(num1)
	fmt.Println(num2)
	// 结论：
	// 1、使用float来计算，可能导致数据不精确。
	// 2、float64的精度>float32， go语言中，浮点数默认使用的是float64
}
