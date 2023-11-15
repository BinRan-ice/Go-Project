package main

import "fmt"

// 定义切片
func main() {

	arr := [4]int{1, 2, 3, 4} // 定长
	fmt.Println(arr)

	var s1 []int // 变长，长度是可变的
	fmt.Println(s1)
	// 切片的空判断，初始的切片中，默认是 nil
	if s1 == nil {
		fmt.Println("切片是空的")
	}

	s2 := []int{1, 2, 3, 4} // 切片 变长
	//
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(s2)
	fmt.Printf("%T,%T\n", arr, s2) // [4]int,[]int
	fmt.Println(s2[1])
}
