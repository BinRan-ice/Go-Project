package main

import "fmt"

// 指针的使用
/*
1、定义指针变量
2、为指针变量赋值  &
3、访问指针变量中地址所指向的值   *
* ：在指针类型前面加上 * 号，就是来获取指针所指向的地址的值

*/
func main() {

	// 声明 普通变量
	var a int = 10
	fmt.Printf("a 变量的值：%d\n", a)   // 10
	fmt.Printf("a 变量的地址：%p\n", &a) // 0x

	// 声明 指针变量，指向a， 指针其实就是一个特殊的变量而已。，ptr命名  p
	// 定义变量格式  var ptr *类型
	var p *int
	p = &a // 指针变量赋值

	fmt.Printf("p 变量存储的指针地址：%p\n", p)     // a 的地址
	fmt.Printf("p 变量自己的地址：%p\n", &p)      // p 变量自己的地址
	fmt.Printf("p 变量存储的指针地址指向的值%d\n", *p) // a的值

	// 改变a的值   *p 和 a 其实是同一个地址
	a = 20
	fmt.Printf("a:%d\n", a)
	fmt.Printf("*p的值：%d\n", *p)

	// 通过指针改变a的值
	*p = 40
	fmt.Printf("a:%d\n", a)
	fmt.Printf("*p的值：%d\n", *p)

	// 指针的套娃，指针指向指针 , 指针类型 第一个*指针类型， *int是这个指针对应的类型
	// 如何理解多个符号，第一个取出来后，后面就是它的类型 *(*(int))
	var ptr **int
	ptr = &p
	//
	fmt.Printf("ptr变量存储的指针的地址：%p\n", ptr)
	fmt.Printf("ptr变量自己的地址：%p\n", &ptr)
	fmt.Printf("*ptr变量存储的地址：%p\n", *ptr)
	fmt.Printf("*ptr变量存储的地址中的值：%d\n", **ptr)
	// 修改变量a就有了无数种方式
	**ptr = 1111
	fmt.Println(a)

}
