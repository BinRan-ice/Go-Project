package main

import "fmt"

func main() {

	// make()
	// make([]Type,length,capacity) // 创建一个切片，长度，容量
	s1 := make([]int, 5, 10)
	fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))
	// 思考：容量为10，长度为5，我能存放6个数据吗？
	s1[0] = 10
	s1[7] = 200 // index out of range [7] with length 5
	// 切片的底层还是数组 [0 0 0 0 0] [2000]
	// 直接去赋值是不行的，不用用惯性思维思考
	fmt.Println(s1)

	// 切片扩容

}
