package pojo

type User struct { // 可以被导出的
	Name  string // 可以被导出的
	money int    // 不可以被导出的
	Age   int    // 可以被导出的
}
