package main

import "fmt"

func main() {
	s1 := make([]int, 0, 5)
	fmt.Println(s1)
	// 切片扩容，append()
	s1 = append(s1, 1, 2)
	fmt.Println(s1)
	// 问题：容量只有5个，那能放超过5个的吗？ 可以，切片是会自动扩容的。
	s1 = append(s1, 2, 3, 3, 4, 4, 5, 6, 6, 7, 7)
	fmt.Println(s1)

	// 切片扩容之引入另一个切片。
	// new ： 解构   slice.. ，解出这个切片中的所有元素。
	s2 := []int{100, 200, 300, 400}
	// slice = append(slice, anotherSlice...)
	// ... 可变参数 ...xxx
	// [...] 根据长度变化数组的大小定义
	// anotherSlice... , slice...解构，可以直接获取到slice中的所有元素
	// s2... = {100,200,300,400}
	s1 = append(s1, s2...)
	// 遍历切片
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}

	for i := range s1 {
		fmt.Println(s1[i])
	}
}
