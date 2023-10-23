package main

import "fmt"

// 全局变量：在当前go文件中生效...
// 定义在go文件非函数内，在package和import下面
// 全局变量的定义必须使用 var 关键字, 如果直接使用 := 则无法创建该变量
// 全局变量和局部变量是可以重名的，优先级。到底用谁
var c int

func main() {
	// 局部变量：只在一定的范围内生效...
	// 在函数体内声明变量
	var a int = 3
	var b int = 4
	// 如果全局变量有，那么直接使用全局变量来接收。
	c = a + b
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
	fmt.Printf("c内存地址:%p\n", &c)
	b = 1
	// 但是如果在局部有和全局同名的变量，优先使用局部变量
	c := a + b
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
	fmt.Printf("c内存地址:%p\n", &c)
	b = 5
	// 就近原则
	c = a + b
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
	fmt.Printf("c内存地址:%p\n", &c)

	// Printf 格式化输出 (参数一：需要打印的内容，%是占位符，通过后续参数给他一一赋值)
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
}
