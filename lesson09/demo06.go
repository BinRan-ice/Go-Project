package main

import "fmt"

// 通过switch来判断  switch i.(T)

type I interface{}

// 如果断言的类型同时实现了switch 多个case匹配，默认使用第一个case
func testAsserts(i interface{}) {
	// switch i.(type) 接口断言
	switch i.(type) {

	case string:
		fmt.Println("变量为string类型")
	case int:
		fmt.Println("变量为int类型")
	case I:
		fmt.Println("变量为I类型")
	case nil:
		fmt.Println("变量为nil类型")
	case interface{}:
		fmt.Println("变量为interface{}类型")

	// .....
	default:
		fmt.Println("未知类型")
	}
}

func main() {

	testAsserts("string")
	testAsserts(1)
	var i I      // 默认值为 nil
	var i2 I = 1 // 只有赋值了之后，才是对应的类型

	testAsserts(i)
	testAsserts(i2)

}
