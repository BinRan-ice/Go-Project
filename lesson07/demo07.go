package main

import "fmt"

// 定一个结构体 type User struct
type User struct {
	name string
	age  int
	sex  string
}

func main() {
	// 通过结构体创建对象， 以前的类型都是用的基本类型，自己定义类型了type，Struct
	// 定义了结构体对象，不赋值，默认都是这个结构体的零值 {"",0,""}
	var user1 User
	fmt.Println("user:", user1)
	// 给结构体对象赋值。 xxx.属性 = 值
	user1.name = "张三"
	user1.age = 1
	user1.sex = "女"
	fmt.Println("user:", user1)

	// 获取这个人的名字
	fmt.Println("user:", user1.name)

	// 创建对象的方式二
	user2 := User{}
	user2.name = "qinjiang"
	user2.age = 27
	user2.sex = "男"
	fmt.Println("user:", user2)
	fmt.Println("user:", user2.name)

	// 	// 创建对象的方式三
	user3 := User{
		name: "feige",
		age:  35,
		sex:  "男",
	}
	fmt.Println("user:", user3)
	fmt.Println("user:", user3.name)

	// 创建对象的方式四，这种不声明属性的方式，需要参数一一匹配
	user4 := User{"guoguo", 18, "女"}
	fmt.Println("user:", user4)
}
