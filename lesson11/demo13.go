package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 读取文件
	file, _ := os.OpenFile("D:\\Environment\\GoWorks\\src\\xuego\\lesson11\\a.txt",
		os.O_RDWR, os.ModePerm)
	// defer close
	defer file.Close()

	// 测试seek
	// 相对开始位置。io.SeekStart
	// 相对于文件末尾， io.SeekEnd
	file.Seek(2, io.SeekStart)
	buf := []byte{0}
	file.Read(buf)

	fmt.Println(string(buf))

	// 相对于当前位置
	file.Seek(3, io.SeekCurrent)
	file.Read(buf)

	fmt.Println(string(buf))

	// 在结尾追加内容
	file.Seek(0, io.SeekEnd)
	file.WriteString("hahahaha")
}
