package main

import "fmt"

// 探究for循环   init：起始值  end：结束值  condition：条件变量
func main() {
	i := 0
	// for 循环可以直接
	// 在for循环中定义的变量，作用域就只在 for循环内，外面是无法调用的
	// for + 结束for条件就可以跑起来
	for i <= 5 {
		i++
	}
	fmt.Println(i)
}
