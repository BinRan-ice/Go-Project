package main

import "fmt"

// 浮点数 float ，默认是 float64 ，32
func main() {
	// 浮点数是不精确的,计算机底层导致的 0 1、浮点数、复数，没办法做到精确计算
	// 浮点数 ； 符号位  + 指数位 + 尾数位（可能会丢失，造成精度损失）。
	var f1 float64
	f1 = 3.16
	var f2 float32
	f2 = 5
	fmt.Println(f1)
	// // 默认的 float 类型，都是保留了6位小数
	// 浮点数的格式化输出 %f ，都是保留了6位小数,保留2位或者3位怎么实现
	// .3 保留3位小数 其余类推
	// 如果小于了当前保留的位数，四舍五入
	fmt.Printf("%T,%.1f\n", f1, f1)
	fmt.Printf("%T,%f", f2, f2)

}
