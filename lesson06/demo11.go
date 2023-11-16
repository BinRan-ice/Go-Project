package main

import "fmt"

// 内存中的一个简单的数据库。key - value 。 增删改查的。

// map的使用
// 定义  var m1 map[key]value    make(map[key]value)
func main() {
	// 创建map
	var map1 map[int]string // nil，不能使用的，只是声明了，还没有初始化
	// 给map1赋值
	// 初始化map
	map1 = make(map[int]string)
	// 在map中如果你的key重复了，它就会覆盖这个key原来的值，一个key只能对应一个value
	map1[100] = "学习"
	map1[100] = "知识"
	map1[200] = "冰冉"

	// 获取数据  map[key]
	fmt.Println(map1)
	fmt.Println(map1[200])
	fmt.Println(map1[1]) // 不存在，默认值 string ""

	// map中，没有index下标，有时候我们取值就需要判断这个key是否存在了
	// map中的判断，ok-idiom 是否存在的
	// value = map[key] , 隐藏的返回值 ok-idiom , 可选参数
	value, ok := map1[1]
	//
	if ok {
		fmt.Println("map key 存在的，value:", value)
	} else {
		fmt.Println("map key 不存在的")
	}

	// 修改数据
	map1[100] = "next"
	fmt.Println(map1)
	// 如果数据存在，修改它，如果不存在，就会创建对应的 key-value
	map1[1] = "xxxxxxxxxxxxxxx"
	fmt.Println(map1)

	// map中的数据删除问题  delete
	delete(map1, 1)
	fmt.Println(map1)

	// map的大小
	fmt.Println(len(map1))

}
