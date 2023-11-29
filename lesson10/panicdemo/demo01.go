package panicdemo

import (
	"fmt"
)

// 错误 error 是一个接口
//
//	type error interface {
//		Error() string
//	}
//
// 自定义一些我们的错误
type KuangShenError struct {
	msg  string
	code int
	// 错误信息是可以定义很多的，主要是描述这个错误出现的具体问题。
	// 方便排错
	data []string
}

func (e *KuangShenError) Error() string {
	return fmt.Sprintf("错误信息:%s\n,代码：%d\n", e.msg, e.code)
}

// error的使用,如果函数或者方法中有预期的错误，都需要抛出。
// func testErr(i int) (int,error)
func testErr(i int) (int, error) {
	if i != 0 {
		return 0, &KuangShenError{msg: "错误数据", code: 500}
	}
	return i, nil
}

func demo01() {
	// 假设这个方法返回了多个错误
	i, err := testErr(1)
	if err != nil {
		fmt.Println(err)
		e, ok := err.(*KuangShenError)
		if ok {
			fmt.Println(e.msg)
		}
	}
	fmt.Println(i)
}
