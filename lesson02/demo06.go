package main

import "fmt"

func main() {

	a := false

	// 这里的爆红不影响  if{....}
	switch a {
	case false: // 基本业务代码
		fmt.Println("1") // 业务代码
		fallthrough      // 在case中 一旦使用了 fallthrough，则会强制执行下一个case语句
	case true: // 善后代码
		fmt.Println("2")
		fallthrough
	case false:
		fmt.Println("3")
		fallthrough
	case true:
		// 判断了一些错误... 跳出终止这个case
		if a == false {
			break
		}
		fmt.Println("4")
		fallthrough
	case false:
		fmt.Println("5")

	default: // else
		fmt.Println("6")
	}

	// 如果满足，则返回基本查询结果，和一个状态结果
	// name  success 用户信息  ok 200... 状态码结果.. 其他东西

}
