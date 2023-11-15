package main

import "fmt"

// 切片扩容的内存分析
// 结论
// 1、每个切片引用了一个底层的数组
// 2、切片本身不存储任何数据，都是底层的数组来存储的，所以修改了切片也就是修改了这个数组中的数据
// 3、向切片中添加数据的时候，如果没有超过容量，直接添加，如果超过了这个容量，就会自动扩容，成倍的增加， copy
//   - 分析程序的原理
//   - 看源码
//
// 4、切片一旦扩容，就是重新指向一个新的底层数组。
func main() {
	// 1、cap 是每次成倍增加的
	// 2、只要容量扩容了，地址就会发生变化
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1)) // len:3,cap:3
	fmt.Printf("%p\n", s1)                          // 0xc000016108

	s1 = append(s1, 4, 5)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1)) // len:5,cap:6
	fmt.Printf("%p\n", s1)                          // 0xc000010390

	s1 = append(s1, 6, 7, 8)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1)) // len:8,cap:12
	fmt.Printf("%p\n", s1)                          // 0xc00005e060

	s1 = append(s1, 9, 10)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1)) // len:10,cap:12
	fmt.Printf("%p\n", s1)                          // 0xc00005e060

	s1 = append(s1, 11, 12, 13, 14)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1)) // len:14,cap:24
	fmt.Printf("%p\n", s1)                          // 0xc00010c000
}
