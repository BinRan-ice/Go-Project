package main

import "fmt"

// 特殊的常量 iota 常量计数器，const （多个常量，自动帮我们 + 1）
func main() {
	// iota 每次定义新的常量，在一组 const 中，那么它会自动 + 1
	// iota 默认值0,
	const (
		a = iota
		b = iota
		c = iota
		d = 0
		e = iota
		// 如果在定义 const 的时候，如果下面的常量没有赋值，默认沿用上面一个常量定义的赋值。
		f
		g
		h = iota
		// ... 1000+
	)
	const (
		i = iota // 0
		j = 0    // j =0  iota + 1 = 1
		k = iota // iota + 1 = 2
	)
	// 0 1 2  d=0 ，e = 4
	fmt.Println(a, b, c, d, e, f, g, h)
	fmt.Println(i, j, k)

}
