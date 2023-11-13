package main

import "fmt"

func main() {
	// 成绩
	var score int = 90

	// 满足一个条件，则进入对应的处理代码
	if score == 100 {
		fmt.Println("满分")
	} else { // 不满足就进入else
		fmt.Println("不及格")
	}

}
