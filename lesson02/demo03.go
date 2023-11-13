package main

import "fmt"

// 命令行程序 go build xxx.go 生成可执行程序。
func main() {

	// 练习：if的练习

	// 判断用户密码输入是否正确

	// 输入框 ： 接收用户的输入 （新知识）
	// 第一次判断
	// 输入框 ：请再次输入密码  接收用户的输入 （新知识）
	// 第二次判断
	// 如果两次都是对的，那么则登录成功

	//
	var num1 int
	var num2 int

	// 提示用户输入
	fmt.Printf("请输入密码 ： \n")
	// 接收用户的输入 （阻塞式等待... 直到用户输入之后才继续运行）最简单的人机交互
	// Scan()  &num1地址，
	fmt.Scan(&num1) // 等待你的输入，卡住... 如果你输入完毕，回车。输入内容就会赋值给num
	// 123456 模拟数据，未来是在数据库中查询出来的。根据用户查的
	if num1 == 123456 {
		fmt.Println("请再次输入密码: ")
		fmt.Scan(&num2)
		if num2 == 123456 {
			fmt.Println("登录成功")
		} else {
			fmt.Println("登录失败")
		}
	} else {
		fmt.Println("登录失败")
	}

}
