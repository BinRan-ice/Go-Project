package main

import (
	"fmt"
	"runtime"
	"time"
)

// 获取系统的信息runtime
func main() {

	// 终止程序 runtime.Goexit()
	go func() {
		fmt.Println("start")
		runtimetest()
		fmt.Println("end")
	}()

	time.Sleep(time.Second * 3)
}
func runtimetest() {
	defer fmt.Println("test defer")
	//return // 只是终止了函数
	runtime.Goexit() // 终止当前的 goroutine
	fmt.Println("test")
}

func runtime1() {
	// 获取goRoot目录 : 找到指定目录，存放一些项目信息。
	fmt.Println("GoRoot Path:", runtime.GOROOT())
	// 获取操作系统  windows ，判断盘符字符。 “\\”  “/”
	fmt.Println("System:", runtime.GOOS)
	// 获取cpu数量 8， 可以尝试做一些系统优化，开启更大的栈空间。
	fmt.Println("Cpu num:", runtime.NumCPU())
}

func runtime2() {

	// goroutine是竞争cpu的  ，调度
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine", i)
		}
	}()

	for i := 0; i < 5; i++ {
		// gosched:礼让, 让出时间片，让其他的 goroutine 先执行
		// cpu是随机，相对来说，可以让一下，但是不一定能够成功
		// schedule
		runtime.Gosched()
		fmt.Println("main-", i)
	}
}
