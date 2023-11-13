package main

import "fmt"

func main() {

	// 分数
	var score int
	fmt.Println("请输入你的成绩：")
	fmt.Scan(&score)
	// 逻辑运算符应用   90<=score<=100 不能这样写的
	if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else if score >= 80 && score < 90 {
		fmt.Println("B")
	} else if score >= 70 && score < 80 {
		fmt.Println("C")
	} else if score >= 60 && score < 70 {
		fmt.Println("D")
	} else if score < 0 || score > 100 { // 健壮性判断
		fmt.Println("输入不合法")
	} else { // else 如果以上的条件都不满足，则执行 else
		fmt.Println("不及格")
	}

}
