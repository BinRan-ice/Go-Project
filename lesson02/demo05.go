package main

import "fmt"

func main() {

	var score int = 100

	// 通过switch来判断score
	// case , 后面可以写多个条件
	switch score {
	case 100, 95, 91:
		fmt.Println("A")
	case 90:
		fmt.Println("B")
	case 80, 70, 60:
		fmt.Println("C")
	default:
		fmt.Println("other")
	}

	// switch 是可以省略条件的，默认是 switch true
	switch {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	default:
		fmt.Println("其他")
	}

	//var flag bool
	//switch flag {
	//case score<100 :
	//
	//}
}
