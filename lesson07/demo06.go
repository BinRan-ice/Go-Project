package main

import "fmt"

func main() {
	// 值传递
	a := 10
	fmt.Println("a:", a)       // 10
	fmt.Println("a addr:", &a) // 0xc00000e0a8

	f6(&a)

	fmt.Println("a:", a)       // 20
	fmt.Println("a addr:", &a) // 0xc00000e0a8
}

// 指针当做函数的参数
func f6(ptr *int) {
	fmt.Println("ptr:", ptr)                  // 0xc00000e0a8
	fmt.Println("ptr 指针的地址中的值", *ptr) // a = 10
	*ptr = 20
}