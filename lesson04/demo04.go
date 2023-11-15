package main

import "fmt"

// 数组是值类型： 所有的赋值后的对象修改值后不影响原来的对象。
func main() {
	//数组类型的样子 [size]type
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [5]string{"kuangshen", "xuexiangban"}
	fmt.Printf("%T\n", arr1) // [4]int
	fmt.Printf("%T\n", arr2) // [5]string

	// 数组的值传递和int等基本类型一致
	arr3 := arr1
	fmt.Println(arr1)
	fmt.Println(arr3)

	// 修改arr3观察arr1是否会变化
	arr3[0] = 12
	fmt.Println(arr1)
	fmt.Println(arr3) // 数组是值传递，拷贝一个新的内存空间
}
