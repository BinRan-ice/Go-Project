package main

import "fmt"

func main() {
	a, b := swap("feige", "kuangshen")
	fmt.Println(a, " ", b)
}

// 返回多个返回值的函数定义
// 案例：交换两个string
// 有多个返回值的情况下，返回值用括号括起来
func swap(x string, y string) (string, string) {
	return y, x
}
