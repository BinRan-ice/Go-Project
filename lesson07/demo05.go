package main

import "fmt"

// 指针函数， 指针是可以用作函数的返回值
func main() {
	// 调用了这个函数后，可以得到一个指针类型的变量。
	ptr := f5()

	fmt.Println("ptr:", ptr)
	fmt.Printf("ptr类型:%T\n", ptr)
	fmt.Println("ptr的地址:\n", &ptr)
	fmt.Println("ptr地址中的值:\n", *ptr)

	// 使用
	fmt.Println((*ptr)[0])
	ptr[0] = 10
	fmt.Println(ptr[0])

}

// 调用该函数后返回一个指针
func f5() *[4]int {
	arr := [4]int{1, 2, 3, 4}
	return &arr
}
