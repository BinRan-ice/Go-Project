package main

import (
	"fmt"
	"time"
)

func main() {
	// 加  减    比较（在xxx之前  在xxx之后  相等）
	now := time.Now()
	later := now.Add(time.Hour)
	fmt.Println(later)

	// 两个时间的差值
	subTime := later.Sub(now)
	fmt.Println(subTime) // 1h0m0s

	// 比较时间, init() 校验时间 当地时间和网络时间是否一致
	fmt.Println(now.Equal(later))  // fasle
	fmt.Println(now.Before(later)) // true
	fmt.Println(now.After(later))  // fasle
}

// 定时器 - 本质是一个通道chan
func d1() {
	// 定时器： 每隔xxx s执行一次, 其余的关于定时器，放到chan讲
	ticker := time.Tick(time.Second) // 每一秒都会触发。
	for i := range ticker {
		fmt.Println(i)
	}
}
