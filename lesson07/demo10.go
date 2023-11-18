package main

import "fmt"

// 一个结构体可能包含一个字段，而这个字段又是一个结构体：结构体嵌套
type Person struct {
	name    string
	age     int
	address Address
}
type Address struct {
	city, state string
}

// 结构体是可以嵌套：我们就可以定义很多复杂的对象，来进行拼接了，构成一个更大的对象
func main() {

	var person = Person{}
	person.name = "binran"
	person.age = 18
	person.address = Address{
		city:  "广州",
		state: "中国",
	}
	fmt.Println(person.name)
	fmt.Println(person.age)
	fmt.Println(person.address)
	fmt.Println(person.address.city)

}
