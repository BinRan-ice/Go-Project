package main

import (
	"fmt"
	"net"
)

// error 是一个类型，error 是一个接口，里面只有一个方法 Error()
// type error interface {
//	 Error() string
// }

func main() {

	// 网络连接
	// 返回这个网址的服务地址（cdn、nginx）
	// func LookupHost(host string) (addrs []string, err error)
	addrs, err := net.LookupHost("www.baidu222222.com")
	if err != nil {
		fmt.Println(err)
	}
	// 如果知道它是什么错误类型的话，这里就可以通过断言来操作
	dnsErr, ok := err.(*net.DNSError)
	if ok {
		if dnsErr.Timeout() {
			// 超时的错误逻辑
			fmt.Println("超时....")
		} else if dnsErr.Temporary() {
			// 临时错误的逻辑
			fmt.Println("临时错误....")
		} else {
			fmt.Println("其他错误....")
		}
	}
	fmt.Println(addrs)
}
