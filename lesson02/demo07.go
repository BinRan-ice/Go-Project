package main

import "fmt"

// 输出0-10，每输出完一个换行
func main() {
	// 了解for循环
	//for 参数1：从哪里开始  参数2：到哪里结束 参数3：控制这个循环条件变量（自增和自减）
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}
