package main

import "fmt"

// 计算1 + 10 的和
func main() {
	sum := 0 // 相加的一个结果

	// goland的快捷输入  fori
	for i := 1; i <= 123133; i++ {
		fmt.Println("i=", i, "  ,  sum=", sum)
		sum = sum + i
	}
	fmt.Println(sum) // 7580929411
}
