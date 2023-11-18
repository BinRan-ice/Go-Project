package main

import "fmt"

// 定一个结构体 type User struct
type User2 struct {
	name string
	age  int
	sex  string
}

func main() {

	// 结构体类型    包.struct名
	user1 := User2{"binran", 18, "男"}
	fmt.Println(user1)
	fmt.Printf("%T,%p\n", user1, &user1) // main.User2,0xc00007e4b0

	// 结构体是值类型的
	user2 := user1
	fmt.Println(user2)
	fmt.Printf("%T,%p\n", user2, &user2) // main.User2,0xc00007e540

	user2.name = "tywin"
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println("========================")
	// 指针解决值传递的问题
	var user_ptr *User2
	user_ptr = &user1
	// *user_ptr 等价于 user1
	fmt.Println(*user_ptr)
	(*user_ptr).name = "qinjiang"
	fmt.Println(user1)
	// 语法糖
	user_ptr.name = "qinjiang222222222"
	fmt.Println(user1)

	// 内置函数 new 创建对象。  new 关键字创建的对象，都返回指针，而不是结构体对象。
	// func new(Type) *Type
	// 通过这种方式创建的结构体对象更加灵活，突破了结构体是值类型的限制。
	user3 := new(User2)
	fmt.Println(user3)

	(*user3).name = "小红"
	user3.sex = "女"
	user3.age = 18
	fmt.Println(user3)

	updateUser(user3)
	fmt.Println(user3)
}
func updateUser(user *User2) {
	user.age = 100
}
