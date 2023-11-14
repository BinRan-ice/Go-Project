package main

import "fmt"

func main() {
	fmt.Println(getSum())
}

// 可变参数： 一个函数的参数类型确定，参数的个数不确定，可以使用可变参数
// 可变参数如果有多个参数必须放在最后一个参数

// 求和 , 参数是可变的，int
func getSum(nums ...int) int {
	sum := 0
	fmt.Println("可变参数的长度为：", len(nums))
	for i := 0; i < len(nums); i++ {
		fmt.Println("可变参数", i, ":", nums[i])
		// 取出来
		//sum = sum + nums[i]
		sum += nums[i]
	}
	return sum
}

// 接收多个参数 nums 可变参数
// 使用下标来接收，下标是从0开始的
// nums : 100,200,300,400,500
// nums[0] = 100
// nums[1] = 200
// nums[2] = 300

// 了解传递了多少个数字  len()函数，获取可变参数的长度
