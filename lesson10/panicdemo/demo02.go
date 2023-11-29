package panicdemo

import "fmt"

// panic  recover
//
//	出现了panic之后，如果有defer语句，先执行所有的defer语句。
//
// defer : 延迟函数，倒序执行，处理一些问题。
func demo02() {
	defer fmt.Println("main--1")
	defer fmt.Println("main--2")
	fmt.Println("main--3")
	testPanic(1) // 外部函数，如果在函数内部已经处理panic,那么程序会继续执行
	defer fmt.Println("main--4")
	fmt.Println("main--5")
}
func testPanic(num int) {
	// 出去函数的时候处理这里面可能发生的panic
	// recover func recover() any 返回panic传递的值
	// panic   func panic(v any)
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("recover执行了... panic msg:", msg)
			// 处理逻辑
			fmt.Println("---------程序已恢复----------")
		}
	}()
	defer fmt.Println("testPanic--1")
	defer fmt.Println("testPanic--2")
	fmt.Println("testPanic--3")
	// 如果在函数中一旦触发了 panic，会终止后面要执行的代码。
	// 如果存在defer，正常按照defer逻辑执行
	if num == 1 {
		panic("出现预定的异常----panic")
	}
	defer fmt.Println("testPanic--4")
	fmt.Println("testPanic--5")
}
