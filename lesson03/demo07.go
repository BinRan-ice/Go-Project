package main

import "fmt"

/*
定义：一个函数自己调用自己，就叫做递归函数
注意：递归函数需要有一个出口，逐渐向出口靠近，没有出口就会形成死循环。
*/
func main() {
	// overflows 栈溢出
	sum := getSum2(1000000000)
	fmt.Println(sum)
}
func getSum2(n int) int {
	fmt.Println(n)
	if n == 1 {
		return 1
	}
	return getSum2(n-1) + n
}

// 假设 getSum(5)

// 求和 sum  1 + 2 + 3 + 4 + 5
// getSum(5)☟=15
//      getSum(4)☟=10 + 5
//          getSum(3)☟=6 + 4
//               getSum(2)☟=3 + 3
// //                getSum(1)=1 + 2
