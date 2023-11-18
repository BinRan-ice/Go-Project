package main

import "fmt"

// 数组指针
func main() {

	// 创建数组，值传递。fun
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("arr1:", arr1)
	fmt.Printf("arr1指向的地址:%p\n", &arr1)

	// 创建一个指针，指向这个数组的地址，通过指针来操作数组
	var p1 *[4]int
	p1 = &arr1
	fmt.Printf("p1指向的地址: %p\n", p1)
	fmt.Printf("p1自己的地址: %p\n", &p1)
	fmt.Println("p1指向的地址的值: ", *p1)

	// 操作数组指针 来修改数组
	(*p1)[0] = 100 // 原生写法
	fmt.Println("arr1:", arr1)
	fmt.Println("p1指向的地址的值: ", *p1)

	// 语法糖：由于p1指向了arr1这个数组，所以可以直接用p1来操控数组
	// 指向了谁，这个指针就可以代表谁。
	// p1 = arr1
	p1[0] = 200 // 在程序中，我们更多时候是这样在使用指针的
	fmt.Println("arr1:", arr1)
	fmt.Println("p1指向的地址的值: ", *p1)
	
}
