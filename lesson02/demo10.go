package main

import (
	"fmt"
)

// 探究for循环   init：起始值  end：结束值  condition：条件变量
func main() {
	i := 0
	// 没有起始值、没有结束...  无限循环，死循环, 不要在开发中写这种代码。会内存溢出
	for {
		i++
		fmt.Println("hello,world")
	}
}
