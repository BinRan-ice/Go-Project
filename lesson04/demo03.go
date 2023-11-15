package main

import "fmt"

/*
1、直接通过下标获取元素 arr[index]

2、 0-len i++ 可以使用for循环来结合数组下标进行遍历

3、for range：范围   (new)
*/
func main() {

	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1[0])
	fmt.Println(arr1[1])
	fmt.Println(arr1[2])
	fmt.Println(arr1[3])
	fmt.Println(arr1[4])
	// 错误：index 5 out of bounds [0:5] 数组下标越界
	// 数组的长度只有5，你要取出6个元素，不可能取出
	//fmt.Println(arr1[5])
	fmt.Println("------------------")
	// 获取数组的长度  len()
	// 下标从0开始，不能<=
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	fmt.Println("------------------")
	// goland 快捷方式 数组.for，未来循环数组、切片很多时候都使用for    range
	// for 下标,下标对应的值  range 目标数组切片
	// 就是将数组进行自动迭代。返回两个值 index、value
	// 注意点，如果只接收一个值，这个时候返回的是数组的下标
	// 注意点，如果只接收两个值，这个时候返回的是数组的下标和下标对应的值
	for _, value := range arr1 {
		fmt.Println(value)
	}

}
