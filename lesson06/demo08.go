package main

import "fmt"

// 切片实现深拷贝
func main() {
	// 将原来切片中的数据拷贝到新切片中
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0) // len:0 cap:0
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)

	// copy
	s3 := []int{5, 6, 7}
	fmt.Println(s2)
	fmt.Println(s3)

	// 将s3中的元素拷贝到s2中
	//copy(s2, s3)
	// 将s2中的元素拷贝到s3中
	copy(s3, s2)
	fmt.Println(s2)
	fmt.Println(s3)
}
