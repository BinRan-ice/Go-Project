package main

import "fmt"

func main() {

	a := 1
	b := 2
	c := 3
	d := 4

	// 创建一个指针数组
	arr1 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr1)

	// 通过指针修改a的值
	// arr1[0] 0xc00000e0a8
	*arr1[0] = 100
	fmt.Println(a)

	a = 200
	fmt.Println(*arr1[0])
}
