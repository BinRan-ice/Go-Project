package main

import (
	"fmt"
	"math/rand"
	"time"
)

// random随机数- math/rand
func main() {
	// 获取一个随机数
	num1 := rand.Int()
	fmt.Println("num1:", num1) // 5577006791947779410
	// 随机需要一个随机数的种子，如果种子一样，那么结果一致
	// n范围（0-n）
	num2 := rand.Intn(100)
	fmt.Println("num2:", num2) // 7

	// 需要一个随时都在发生变化的量 时间戳
	timestamp := time.Now().Unix()
	// 设置随机数种子, 使用时间戳
	// 种子只需要设置一次即可。
	rand.Seed(timestamp) // 每次执行都不同
	for i := 0; i < 5; i++ {
		// Intn [0,n)
		// 20-29
		num1 := rand.Intn(200) // 抽奖程序
		// 必中的逻辑
		num2 := rand.Intn(5) // 抽奖程序
		if i == num2 {
			fmt.Println(10)
			continue
		}
		fmt.Println(num1)
	}

}
