package main

import (
	"fmt"
	"os"
)

// IO读
func main() {

	// 找到这个文件的对象 create 创建、 打开Open
	// func Open(name string) (*File, error)
	file1, err := os.Open("D:\\Environment\\GoWorks\\src\\xuego\\lesson11\\狂神的Go世界.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file1)
	// file 类-- 指定的对象
	// 打开文件的时候，选定权限： 可读可写的方式打开
	// OpenFile（文件名，打开方式：可读、可写...，FileMode ， 权限）
	file2, err2 := os.OpenFile("D:\\Environment\\GoWorks\\src\\xuego\\lesson11\\狂神的Go世界.txt",
		os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(file2)

	// 可以操作这个对象了

}
