package main

import "fmt"

func main() {
	// var 变量名 变量类型
	// 定义变量 ，如果没有给这个变量赋值， 就是这个数据类型的默认值。
	// string 默认值 ”“  int 默认值 0
	var name string
	var age int

	// 可以同时定义多个变量，也只需要使用 var() 关键字
	var (
		addr  string
		phone int
	)

	// goland 快捷键 ： 复制当前行到下一行  Ctrl + D
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(addr)
	fmt.Println(phone)

	// 变量的使用，在定义完毕变量之后直接可以操作这个变量
	// 给变量赋值 ，符号  = 赋值（不能叫等于）
	// 将 "秦疆" 赋值 给 name 这个变量。
	name = "秦疆"
	// 使用变量，直接打印或者进行一些操作都可以
	fmt.Println(name)

	// 变量是可以被重复赋值的，变量。
	name = "张三"
	fmt.Println(name)

	// 在定义变量的时候直接进行赋值。
	var dog string = "旺财"
	fmt.Println(dog)

}
