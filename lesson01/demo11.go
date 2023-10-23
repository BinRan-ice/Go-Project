package main

import "fmt"

func main() {
	// bool类型，只有两个值 true 和 false。  对和错
	// 定义变量  bool 布尔类型的定义 与 int、string 都是Go语言中基本数据类型。
	var b1 bool
	var b2 bool

	b1 = true
	b2 = false
	// f 格式化输出  %d 整数  %s 字符串 %p内存地址  %T 类型 %t bool值
	fmt.Printf("b1=%T,%t\n", b1, b1)
	fmt.Printf("b2=%T,%t\n", b2, b2)

	// 比大小
	var a int = 1
	var b int = 2
	// 如果  xxx  否则 xxx
	// 结果就是bool类型
	fmt.Println(a > b)
	if a > b {
		fmt.Println("a是大于b的")
		// .....
	} else {
		fmt.Println("a是小于b的")
	}

	// bool 类型的默认值 false， 规定  false  0   true = 1
	var b3 bool
	fmt.Println("bool默认值：", b3)

}
