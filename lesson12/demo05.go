package main

import (
	"fmt"
	"sync"
	"time"
)

// waitgroup、

var wg1 sync.WaitGroup

func main() {
	// 公司最后关门的人   0
	// wg.Add(2) 判断还有几个线程、计数  num=2
	// wg.Done() 我告知我已经结束了  -1
	wg1.Add(2)

	go test1()
	go test2()

	fmt.Println("main等待ing")
	wg1.Wait() // 等待 wg 归零，才会继续向下执行
	fmt.Println("end")

	// 理想状态：所有协程执行完毕之后，自动停止。
	//time.Sleep(3 * time.Second)

}
func test1() {
	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println("test1--", i)
	}
	wg1.Done()
}
func test2() {
	defer wg1.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("test2--", i)
	}
}
