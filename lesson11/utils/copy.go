package utils

import (
	"fmt"
	"io"
	"os"
)

// Copy 方法需要参数为：source 源文件 ,destination 目标文件
func Copy(source, destination string, bufferSize int) {

	// 读取文件
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println("Open错误:", err)
	}
	// 输出文件 O_WRONLY , O_CREATE 如果不不存在，则会创建
	destinationFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("OpenFile错误:", err)
	}
	// 关闭
	defer sourceFile.Close()
	defer destinationFile.Close()

	// 专注业务代码，拷贝
	buf := make([]byte, bufferSize)
	// 读取
	for {
		n, err := sourceFile.Read(buf)
		if n == 0 || err == io.EOF {
			fmt.Println("读取完毕源文件,复制完毕")
			break
		} else if err != nil {
			fmt.Println("读取错误:", err)
			return // 错误之后，必须要return终止函数执行。
		}
		// 将缓冲区的东西写出到目标文件
		_, err = destinationFile.Write(buf[:n])
		if err != nil {
			fmt.Println("写出错误：", err)
		}
	}

}

// 调用系统的方法
func Copy2(source, destination string) {
	// 读取文件
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println("Open错误:", err)
	}
	// 输出文件 O_WRONLY , O_CREATE 如果不不存在，则会创建
	destinationFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("OpenFile错误:", err)
	}
	// 关闭
	defer sourceFile.Close()
	defer destinationFile.Close()

	// 具体的实现
	written, err := io.Copy(destinationFile, sourceFile)
	fmt.Println("文件的字节大小:", written)
}

func Copy3(source, destination string) {
	fileBuf, _ := os.ReadFile(source)
	os.WriteFile(destination, fileBuf, 0777)
}
