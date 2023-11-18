package main

import (
	"GoProject/lesson07/base/pojo"
	"fmt"
)

// 程序入口
func main() {
	var user pojo.User
	user.Name = "binran"
	user.Age = 18
	fmt.Println(user)
}
