package main

import "fmt"

// 数组
// 数组和其他变量定义没什么区别，唯一的就是这个是一组数，需要给一个大小  [6]int   [10]string
// 数组是一个相同类型数据的==有序==集合，通过下标来取出对应的数据
// 数组几个特点：
// 1、长度必须是确定的，如果不确定，就不是数组，大小不可以改变
// 2、元素必须是相，同类型不能多个类型混合, [any也是类型，可以存放任意类型的数据]
// 3、数组的中的元素类型，可以是我们学的所有的类型，int、string、float、bool、array、slice、map

func main() {
	// array数组定义，变量
	// 数组也是一个数据类型
	// 数组的定义：  [数组的大小size]变量的类型 ，
	// 我们定义了一组这个类型的数组集合，大小为size，最多可以保存size个数
	var arr1 [5]int
	// [0,0,0,0,0]
	// 给数组赋值,下标index，所有的数组下标都是从0开始的。
	arr1[0] = 100
	arr1[1] = 200
	arr1[2] = 300
	arr1[3] = 400
	arr1[4] = 500

	// 打印数组
	fmt.Println(arr1)

	// 取出数组中的某个元素
	fmt.Println(arr1[1])

	// 数组中的常用方法 len()获取数组的长度  cap() 获取数组的容量
	fmt.Println("数组的长度：", len(arr1))
	fmt.Println("数组的容量：", cap(arr1))

	// 修改数组的值，index 1 代表的第二个数据了
	arr1[1] = 10
	fmt.Println(arr1)
	fmt.Println(arr1[1])

}
