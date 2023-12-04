package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义全局变量 票库存为10张
var ticket int = 10

// 定义一个锁  Mutex 锁头
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	wg.Add(4)

	// 单线程不存在问题，多线程资源争抢就出现了问题
	go saleTickets("张三")
	go saleTickets("李四")
	go saleTickets("王五")
	go saleTickets("赵六")

	wg.Wait()

}

// 售票函数
func saleTickets(name string) {
	defer wg.Done()
	for {
		// 在拿到共享资源之前先上锁
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Millisecond * 1)
			fmt.Println(name, "剩余票的数量为：", ticket)
			ticket--
		} else {
			// 操作完毕后，解锁
			mutex.Unlock()
			fmt.Println("票已售完")
			break
		}
		// 操作完毕后，解锁
		mutex.Unlock()
	}
}
