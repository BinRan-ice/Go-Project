package main

import "fmt"

func main() {
	_, _, _, r4 := fun2(2, 4)
	fmt.Println(r4)
}

// 周长、面积案例  （长方形 ... 长、宽）
// 返回多个值，需要括号，需要表明返回值类型，返回值也可以命名
// return的结果值命名和定义函数返回值的命名无关.
func fun1(len, wid float64) (zc float64, area float64) {
	area = len * wid
	zc = (len + wid) * 2
	fmt.Println("zc:", zc)
	fmt.Println("area:", area)
	// 1、return 如果直接定义了，那么返回结果按照 return 的顺序
	// 2、直接调用return不带结果，那么则返回 函数返回值定义的顺序进行结果返回。
	return
}

func fun2(len, wid float64) (float64, float64, float64, float64) {
	area := len * wid
	zc := (len + wid) * 2
	fmt.Println("zc:", zc)
	fmt.Println("area:", area)
	// 1、return 如果直接定义了，那么返回结果按照 return 的顺序
	// 2、直接调用return不带结果，那么则返回 函数返回值定义的顺序进行结果返回。
	return area, zc, 1, 3
}
