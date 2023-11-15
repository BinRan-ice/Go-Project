package main

import "fmt"

// copy 方法
func main() {
	numbers := []int{1, 2, 3}
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)

	// 方法一： 直接使用make创建切片扩容
	numbers2 := make([]int, len(numbers), cap(numbers)*2)
	// 将原来的底层数据的值拷贝到新的数组中
	// func copy(dst, src []Type) int
	copy(numbers2, numbers)

	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers2), cap(numbers2), numbers2)
}
