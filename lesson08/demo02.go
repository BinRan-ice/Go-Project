package main

import "fmt"

// 方法：可以理解为函数多了一个调用者
// 方法可以重名，不同的对象，调用的结果是不一样的
type Dog1 struct {
	name string
	age  int
}

// 方法定义,   func 方法调用者  方法名()
// 1、方法可以重名，只需要调用者不同
// 2、如果调用者相同，则不能重名
func (dog Dog1) eat() {
	fmt.Println("Dog eating...")
}

func (dog Dog1) sleep() {
	fmt.Println("Dog sleep...")
}

type Cat1 struct {
	name string
	age  int
}

func (cat Cat1) eat() {
	fmt.Println("Cat eating...")
}
func (cat Cat1) sleep() {
	fmt.Println("Cat sleep...")
}

func main() {

	// 创建一个对象
	dog := Dog1{
		name: "旺财",
		age:  2,
	}
	fmt.Println(dog)
	// 方法的调用，通过对应的结构体对象来调用
	dog.eat()

	cat := Cat1{name: "喵喵", age: 1}
	cat.eat()
}
