package main

import "fmt"

func main() {

	// 定义一个多维数组  二维

	arr := [3][4]int{
		{0, 1, 0, 0}, // arr[0]  //数组
		{0, 0, 1, 0}, // arr[1]
		{0, 2, 0, 0}, // arr[2]
	}

	// 二维数组，一维数组存放的是一个数组
	fmt.Println(arr[0])
	// 要获取这个二维数组中的某个值，找到对应一维数组的坐标，arr[0] 当做一个整体
	fmt.Println(arr[0][1])
	fmt.Println("------------------")
	// 如何遍历二维数组
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}
	// for range
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
