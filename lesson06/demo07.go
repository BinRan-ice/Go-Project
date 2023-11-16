package main

import "fmt"

func main() {
	// 值类型，传递：copy
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1
	fmt.Println(arr1, arr2)
	arr1[0] = 100
	fmt.Println(arr1, arr2)

	// 切片：引用类型
	s1 := []int{1, 2, 3, 4}
	s2 := s1
	fmt.Println(s1, s2)
	s1[0] = 100
	fmt.Println(s1, s2)
	// 通过内存分析发现，s1,s2指向了同一个地址。
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
}
