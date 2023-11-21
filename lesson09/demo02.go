package main

import (
	"errors"
	"fmt"
)

// 自己定义一个错误
// 1、errors.New("xxxxx")
// 2、fmt.Errorf()
// 都会返回  error 对象， 本身也是一个类型
func main() {

	age_err := setAge(-1)
	if age_err != nil {
		fmt.Println(age_err)
	}
	fmt.Printf("%T\n", age_err) // *errors.errorString

	// 方式二
	errInfo1 := fmt.Errorf("我是一个错误信息：%d\n", 500)
	//errInfo2 := fmt.Errorf("我是一个错误信息：%d\n", 404)
	if errInfo1 != nil {
		// 处理这个错误
		fmt.Println(errInfo1)
	}

}

// 设置年龄的函数，一定需要处理一些非正常用户的请求
// 返回值为 error 类型
// 作为一个开发需要不断思考的事情，代码的健壮性和安全性
func setAge(age int) error {
	if age < 0 {
		// 给出一个默认值
		age = 3
		// 抛出一个错误 errors 包
		return errors.New("年龄不合法")
	}
	// 程序正常的结果，给这个用户赋值
	fmt.Println("年龄设置成功：age=", age)
	return nil
}
