package main

import "fmt"

func main() {

	// 只定义变量，不使用可以吗. 理论可以，实际在Go语言中不行。
	// 无意义代码，报错！

	// 问题1 ：能不能不写数据类型
	// 问题2 ：不用var 直接定义变量呢？

	// 自动推导，一个短变量声明
	name := "kuangshen"
	age := 18.22
	// 语法糖（偷懒，简化开发！）
	// := 相当于快递定义变量。如果给这个变量赋值了，那么会自动推导它的类型
	// var 、 数据类型的定义都被省略的。
	// 数据类型 go语言中基本的数据类型。

	fmt.Println(name)
	fmt.Println(age)

	// 定义了变量name2
	var name2 string
	// 在快速声明中，如果 := 左边的变量已经存在了，那么它将无法创建，无法重新定义
	name3 := "qinjiang222"
	fmt.Println(name2)
	fmt.Println(name3)
	name3 = "zhangsan666"
}
