package main

import (
	"fmt"
	"time"
)

//
func main() {
	// 临界资源：多个协程共享的变量，会导致程序结果位置
	a := 1

	go func() {
		a = 2
		fmt.Println("goroutine a:", a)
	}()

	a = 3
	time.Sleep(3 * time.Second)
	fmt.Println("main a:", a)
}
