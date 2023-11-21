package main

import (
	"fmt"
)

type BinRanError struct {
	msg  string
	code int
}

// 处理error的逻辑
func (e BinRanError) print() bool {
	fmt.Println("hello，world")
	return true
}

// Error() string
func (e *BinRanError) Error() string {
	//  fmt.Sprintf() 返回string
	return fmt.Sprintf("错误信息：%s,错误代码：%d\n", e.msg, e.code)
}

// 使用错误测试
func test(i int) (int, error) {
	// 需要编写的错误
	if i != 0 {
		// 更多的时候我们会使用自定义类型的错误
		return i, &BinRanError{msg: "非预期数据", code: 500}
	}
	// 正常结果
	return i, nil
}

func main() {

	i, err := test(1)

	if err != nil {
		fmt.Println(err)
		ks_err, ok := err.(*BinRanError)
		if ok {
			if ks_err.print() {
				// 处理这个错误中的子错误的逻辑
			}
			fmt.Println(ks_err.msg)
			fmt.Println(ks_err.code)
		}
	}

	fmt.Println(i)

}
