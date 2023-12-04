package main

import (
	"fmt"
)

// main函数如果结束了，所有的 goroutine也会瞬间销毁
func main() {
	// goroutine ： 和普通方法调用完全不同，它是并发执行的，快速交替。
	go hello()
	fmt.Println("a")
}

// hello函数
func hello() int {
	for i := 0; i < 1000; i++ {
		fmt.Println("hello - ", i)
	}
	return 1
}
