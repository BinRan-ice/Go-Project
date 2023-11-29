package stringsdemo

import (
	"fmt"
	"strings"
)

// strings 字符串常用操作包
// strings 所有方法自己点一下
// 看源码，这个函数如何使用
// 想一个案例测试
func Test() {

	// 1、字符是不能修改的
	str := "xuexiangban,kuangshen"

	// strings下的常用方法
	// 1、判断某个字符是否包含了指定的内容 Contains
	// Contains(s, substr string) bool
	fmt.Println(strings.Contains(str, "z"))

	// 2、判断某个字符串是否包含了多个字符串中的某一个
	fmt.Println(strings.ContainsAny(str, "zk"))

	// 3、统计这个字符在指定字符串中出现的数量 Count() 计数
	// func Count(s, substr string) int
	fmt.Println(strings.Count(str, "n")) // 4

	// 后期作业和IO会结合的
	fileName := "20230219.mp3"
	// 4、判断用什么开头的HasPrefix()
	if strings.HasPrefix(fileName, "2023") {
		fmt.Println("找到2023开头的文件：", fileName)
	}
	// 5、判断用什么结尾的 HasSuffix()
	if strings.HasSuffix(fileName, ".mp4") {
		fmt.Println("找到mp4结尾的文件：", fileName)
	}

	// 6、寻找这个字符串第一次出现的位置 Index()
	fmt.Println(strings.Index(str, "z"))
	// 7、寻找这个字符串最后一次出现的位置 LastIndex()
	fmt.Println(strings.LastIndex(str, "ua"))

	// 8、拼接字符串, 数组或者切片拼接 ，前端给了我们多个参数。保存为一个字符串
	// Join()
	str2 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(strings.Join(str2, " "))
	// 通过某个格式，拆分字符串 Split()
	str3 := strings.Join(str2, "-")
	fmt.Println(strings.Split(str3, "-")) // 用的最多

	// 大小写ToUpper()
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	// 替换 -1 改所有的，  1 就是一个  2就是2两个
	fmt.Println(strings.Replace(str, "e", "狂神", 1))
	// 截取某个字符串
	str5 := str[0:5]
	fmt.Println(str5)

}
