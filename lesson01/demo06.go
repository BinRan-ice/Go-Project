package main

import "fmt"

func main() {
	// 变量到底是个什么玩意？
	// num 实际上是一片内存空间
	// 我们想要看一个变量的内存地址，只需要在变量名前加上 & 即可。
	// 取地址符 （指针）
	var num int
	num = 1000
	// 思考：这个num在计算机中是什么样子的。 num
	fmt.Printf("num的值:%d,内存地址:%p\n", num, &num)
	num = 2000
	fmt.Printf("num的值:%d,内存地址:%p\n", num, &num)

	// 汇编。理解一切

	var name string
	name = "张三"
	// 思考：这个num在计算机中是什么样子的。 num
	fmt.Printf("num的值:%s,内存地址:%p\n", name, &name)
	name = "李四"
	fmt.Printf("num的值:%s,内存地址:%p\n", name, &name)

	// 打印内存地址的方式之一。 Print  f格式化输出
	// 内存
	// 第一个参数 输出字符串
	// % 占位符。
	// 占位符的数量，要和后面待输出的数量一直
	// %d 数字 int d
	// %p 内存地址，num需要取出变量的地址。
	// %s 字符串。
	// \n 换行
	//fmt.Printf("num的值:%d",num)
}
