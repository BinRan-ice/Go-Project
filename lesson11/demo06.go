package main

import (
	"fmt"
	"os"
)

// file
// fileInfo
/*
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits ： 权限
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()

	// 获取更加详细的文件信息， *syscall.Stat_t  反射来获取
	Sys() any           // underlying data source (can return nil)
*/
func main() {
	// 获取某个文件的状态
	fileinfo, err := os.Stat("D:\\Environment\\GoWorks\\src\\xuego\\lesson11\\demo01.go")
	if err != nil {
		return
	}
	fmt.Println(fileinfo.Name())    // demo01.go
	fmt.Println(fileinfo.IsDir())   // false
	fmt.Println(fileinfo.ModTime()) // 2023-02-23 20:25:43.1772351 +0800 CST
	fmt.Println(fileinfo.Size())    // 1186 字节数
	fmt.Println(fileinfo.Mode())    // -rw-rw-rw-
}
