package main

import "fmt"

func main() {
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("arr1:", arr1)
	update(arr1)
	fmt.Println("end arr1:", arr1)

	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1:", s1)
	update2(s1)
	fmt.Println("end s1:", s1)
}

// 函数补充：在使用函数的时候，一定要特别注意参数问题，如果是值类型的，很多传递是无效的。
// 一些值传递的类型的参数，如果我们想通过函数来进行修改对应的值，这个时候就需要使用指针

// 指针变量 -> 指向原来变量的地址

// 数组是值类型的
func update(arr [4]int) {
	fmt.Println("--> arr:", arr)
	arr[0] = 100
	fmt.Println("--> end arr:", arr)
}

// 切片是引用类型的
func update2(s []int) {
	fmt.Println("--> s:", s)
	s[0] = 100
	fmt.Println("--> end s:", s)
}
