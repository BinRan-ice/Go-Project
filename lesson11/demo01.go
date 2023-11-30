package main

import (
	"fmt"
	"strconv"
)

// 项目：前端（网页、小程序、app）  后端代码-接收前端的请求
/*
func http(request) response{
	string:= request.getUrlParm
	// 数据库
	// 计算
	// 字符串的转换
}
*/
// string convert  = strconv
func main() {
	// bool
	s1 := "true"
	// 转化 - 字符串转bool（解析：parse）
	// func ParseBool(str string) (bool, error)
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%t\n", b1, b1) // bool,true
	// bool转字符串（格式化 format）
	s2 := strconv.FormatBool(b1)
	fmt.Printf("%T,%s\n", s2, s2) // string,true

	// 整数->字符串  Format   字符串->整数
	s3 := "100000"
	// 整数： 数字、进制、大小
	// 参数：1、str   2、 进制（10）  3、大小
	i1, _ := strconv.ParseInt(s3, 10, 64)
	fmt.Printf("%T,%d\n", i1, i1) // int64,100

	s4 := strconv.FormatInt(i1, 10)
	fmt.Printf("%T,%s\n", s4, s4) // string,100000

	// 10进制转换字符串，简便方法  atoi  itoa
	atoi, _ := strconv.Atoi("-20")
	fmt.Printf("%T,%d\n", atoi, atoi) // int,-20
	itoa := strconv.Itoa(30)
	fmt.Printf("%T,%s\n", itoa, itoa) // string,30

}
