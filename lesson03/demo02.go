package main

import "fmt"

/*
*
- 无参无返回值
- 有一个参数的函数
- 有两个 or 多个参数的函数
- 有一个返回值的函数
- 有两个 or 多个返回值的函数
*/
func main() {
	m := max(123123123, 12312312312312)
	fmt.Println(m)
}

/*
func 函数名(参数1,参数2) 返回值类型 {
	// 代码逻辑

	return xxx
}
*/
// 定义一个稍微复杂一点的函数
// 比大小的函数 max ，两个数字比大小
// max函数，需要两个参数，int 类型的，比完大小之后我希望返回大的那一个数值
func max(num1 int, num2 int) int {
	var result int
	if num2 > num1 {
		result = num2
	} else {
		result = num1
	}
	// return返回结果
	return result
}
