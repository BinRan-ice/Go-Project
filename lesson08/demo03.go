package main

import (
	"fmt"
)

// 方法重写，建立在父类和子类结构上的
type Animal1 struct {
	name string
	age  int
}

func (animal Animal1) eat() {
	fmt.Println(animal.name, " 正在吃....")
}
func (animal Animal1) sleep() {
	fmt.Println(animal.name, " 正在sleep....")
}

// 父类更通用，子类更具体

// 子类
type Dog2 struct {
	Animal1
}

// 子类自己的方法
func (dog Dog2) wang() {
	fmt.Println("wangwangwanwg~~~")
}

type Cat2 struct {
	Animal1
	color string // 子类可以定义自己的属性
}

// 子类重写父类的方法 , 子类的方法名和父类同名，即可重写父类的方法
func (cat Cat2) eat() {
	fmt.Println(cat.name, " 正在吃cat....")
}

func main() {
	// 定义一个子类，使用父类方法
	dog := Dog2{Animal1{name: "旺财", age: 3}}
	dog.eat()  // 调用父类的方法
	dog.wang() // 调用自己扩展的方法
	cat := Cat2{Animal1{name: "煤球", age: 3}, "黑色"}
	cat.eat() // 如果重写了父类的方法就是调用子类自己的方法

	fmt.Println(cat.color) // 调用子类自己的属性

	// 子类可以操作父类的方法，父类可以操作子类的吗？不可以
	// 父类不能调用子类自己的扩展的方法
	a := Animal1{name: "大黄", age: 3}
	a.eat()
	a.sleep()

}
