package main

import (
	"fmt"
	"os"
)

// 读取文件数据
func main() {
	// 我们习惯于在建立连接时候通过defer来关闭连接，保证程序不会出任何问题，或者忘记关闭
	// 建立连接
	file, _ := os.Open("狂神的Go世界.txt")
	// 关闭连接
	defer file.Close()

	// 读代码 ,Go 的错误机制，让我们专心可以写业务代码。

	// 1、创建一个容器 （二进制文本文件--0100101010 => 读取流到一个容器 => 读取容器的数据）
	bs := make([]byte, 2, 1024) // 缓冲区，可以接受我们读取的数据
	// 2、读取到缓冲区中。 // 汉字一个汉字占 3个位置
	n, err := file.Read(bs)
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println(string(bs)) // 读取到的字符串  ab
	// 光标不停的向下去指向，读取出来的内容就存到我们的容器中。
	file.Read(bs)
	fmt.Println(string(bs)) // 读取到的字符串  cd
	file.Read(bs)
	fmt.Println(string(bs)) // 读取到的字符串  e
	n, err = file.Read(bs)
	fmt.Println(n)
	fmt.Println(err)        // EOF ，读取到了文件末尾。就会返回EOF。
	fmt.Println(string(bs)) // 读取到的字符串
	n, err = file.Read(bs)
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println(string(bs)) // 读取到的字符串

}
