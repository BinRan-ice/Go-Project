package main

import (
	"bufio"
	"fmt"
	"os"
)

// 写入
func main() {
	file, _ := os.OpenFile("D:\\Environment\\GoWorks\\src\\xuego\\lesson11\\a.txt",
		os.O_RDWR|os.O_CREATE,
		os.ModePerm)
	defer file.Close()

	// bufio
	fileWrite := bufio.NewWriter(file)
	writeNum, _ := fileWrite.WriteString("kuangshen")
	fmt.Println("writeNum:", writeNum)
	// 发现并没有写出到文件，是留在了缓冲区，所以我们需要调用 flush 刷新缓冲区
	// 手动刷新进文件
	fileWrite.Flush()
}
