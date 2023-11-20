package main

import "fmt"

type A interface{}
type Dogg struct {
	name string
}
type Catt struct {
	name string
}

func testNow(a A) {
	fmt.Println(a)
}

// 可以指定定义空接口
// // any is an alias for interface{} and is equivalent to interface{} in all ways.
//type any = interface{}
func testNow2(temp interface{}) {
}

func main() {
	var a1 A = Catt{name: "喵喵"}
	var a2 A = Dogg{name: "旺财"}
	var a3 A = 1
	var a4 A = "xuexiangban"
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	// map
	map1 := make(map[string]interface{})
	map1["name"] = "qinjiang"
	map1["age"] = 18
	fmt.Println(map1)

	// slice
	s1 := make([]any, 0, 10)
	s1 = append(s1, 1, "12312", false, a1, a2)
	fmt.Println(s1)

	var arr [4]interface{}
	fmt.Println(arr)

}
