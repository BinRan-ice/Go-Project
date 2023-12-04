package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件  Create
	// func Create(name string) (*File, error) {
	// 返回的file对象就是我们的文件
	file1, err := os.Create("a.go") // 相对路径
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file1)
	// 删除
	os.Remove("D:\\Environment\\GoWorks\\src\\xuego\\lesson11\\a.go")

}
