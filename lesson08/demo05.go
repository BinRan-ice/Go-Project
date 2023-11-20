package main

import "fmt"

type Animal interface {
	eat()
	sleep()
}

type Dog struct {
	name string
}

func (dog Dog) eat() {
	fmt.Println(dog.name, "--eat")
}
func (dog Dog) sleep() {
	fmt.Println(dog.name, "--sleep")
}

// 多态
func main() {

	// Dog 两重身份：1、Dog   2、Animal ，多态
	dog1 := Dog{name: "旺财"}
	dog1.eat()
	dog1.sleep()

	// Dog 也可以是 Animal
	test2(dog1)

	// 定义一个类型可以为接口类型的变量
	// 实际上所有实现类都可以赋值给这个对象
	var animal Animal // 模糊的 -- 具体化，将具体的实现类赋值给他，才有意义
	animal = dog1
	test2(animal)
}

// Animal 接口
func test2(a Animal) {
	a.eat()
	a.sleep()
}
