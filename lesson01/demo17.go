package main

import "fmt"

func main() {
	// 特殊的字符 转义字符 \n
	// \n 换行
	fmt.Println("hello\nworld")
	// \b   backspace  删除上一个字符
	fmt.Println("hello\bworld")
	// \t  Tab
	fmt.Println("hello\tworld")

	// 转义，就是可以将特殊转义符号或者一些符号打印出来 \

	// ` ` 可以通过它来定义一些长字符串，可以换行输入
	fmt.Println(`\\n
sadad
asd
a
sd
asda
sd
asda`)

}
