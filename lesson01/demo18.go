package main

import "fmt"

// 类型转换
func main() {
	// 数据类型转换在Go中就一个格式
	// 新类型的值 = 新类型（旧类型的值）

	// 高位向低位转 （int64  int32  int16）

	// 浮点数转整数，截断，只保留整数部分
	a := 5.9    // float
	b := int(a) // b 就是int类型的a = 5

	fmt.Printf("%T,%.1f\n", a, a)
	fmt.Printf("%T,%d\n", b, b)

	c := 1
	d := float64(c)
	fmt.Printf("%T,%d\n", c, c)
	fmt.Printf("%T,%f\n", d, d)

	// 布尔类型转换，布尔类型 是不支持类型转换的。
	//var flag bool = true
	//f := int(flag)

}
