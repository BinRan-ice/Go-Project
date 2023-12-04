package main

import (
	"fmt"
	"os"
)

func main() {

	fileName := "狂神的Go世界.txt"
	// 权限：如果我们要向一个文件中追加内容， O_APPEND, 如果没有，就是从头开始写
	file, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_RDONLY|os.O_APPEND, os.ModePerm)
	defer file.Close()

	// 操作
	bs := []byte{65, 66, 67, 68, 69} // A B C D E
	n, err := file.Write(bs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	// string类型的写入 ， 爬虫  string[]
	// net + io、就可以实现很多
	n, err = file.WriteString("hhahahahah哈哈哈哈哈哈哈")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
