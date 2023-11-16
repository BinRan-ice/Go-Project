package main

import "fmt"

// 切片本身是不保存数据的，只是指向了底层数组。
func main() {
	// 数组 [0,10)   数组的截取 [start,end]
	// 数组
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 0xc000018230
	// 通过数组来创建切片 [start,end)
	/*
	  arr = 10个数
	  1,  0xc000000001
	  2,  0xc000000002
	  3,  0xc000000003
	  4,  0xc000000004
	  5,  0xc000000005
	*/
	s1 := arr[:5]  // 1-5   // 0xc000018230 {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := arr[3:8] // 4-8   // 0xc000018248
	s3 := arr[5:]  //  6-10 // 0xc000018258
	s4 := arr[:]   //  0-10 // 0xc000018230

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	// 查看容量和长度
	fmt.Printf("s1 len:%d, cap:%d\n", len(s1), cap(s1)) // s1 len:5, cap:10
	fmt.Printf("s2 len:%d, cap:%d\n", len(s2), cap(s2)) // s2 len:5, cap:7
	fmt.Printf("s3 len:%d, cap:%d\n", len(s3), cap(s3)) // s3 len:5, cap:5
	fmt.Printf("s4 len:%d, cap:%d\n", len(s4), cap(s4)) // s4 len:10, cap:10

	// 内存地址
	fmt.Printf("%p\n", s1) // 指向了原数组
	fmt.Printf("%p\n", s2) // 截断了后，指向了最新截断的元素的下标开始
	fmt.Printf("%p\n", s3)
	fmt.Printf("%p\n", &arr[5])
	fmt.Printf("%p\n", s4) // 指向了原数组
	fmt.Printf("%p\n", &arr)

	// 修改数组的内容, 切片也随之发生了变化 （切：切片不保存数据-->底层的数组 ）
	arr[2] = 100
	fmt.Println(arr) // [1 2 100 4 5 6 7 8 9 10]
	fmt.Println(s1)  // [1 2 100 4 5]
	fmt.Println(s2)  // [4 5 6 7 8]
	fmt.Println(s3)  // [6 7 8 9 10]
	fmt.Println(s4)  // [1 2 100 4 5 6 7 8 9 10]
	fmt.Println("------------------------")
	// 修改切片的内容，发现数组也随之发生了变化。（本质：修改的都是底层的数组）
	s2[2] = 80
	fmt.Println(arr) // [1 2 100 4 5 80 7 8 9 10]
	fmt.Println(s1)  // [1 2 100 4 5]
	fmt.Println(s2)  // [4 5 80 7 8]
	fmt.Println(s3)  // [80 7 8 9 10]
	fmt.Println(s4)  // [1 2 100 4 5 80 7 8 9 10]
	fmt.Println("------------------------")
	// 切片扩容
	s1 = append(s1, 11, 12, 13, 14, 15, 16)
	fmt.Println(arr) // [1 2 100 4 5 80 7 8 9 10]
	// 如果在切片容量内加，会导致原来的数组数据发生修改。只有扩容之后，才会导致底层数组会产生拷贝。新的数组
	fmt.Println(s1) // [1 2 100 4 5 11 12 13 14 15 16]
	fmt.Println(s2) // [4 5 80 7 8]
	fmt.Println(s3) // [80 7 8 9 10]
	// 查看内存地址，已经不是同一个了
	fmt.Printf("%p\n", s1)   // 0xc0000d2000
	fmt.Printf("%p\n", &arr) // 0xc0000ac0f0

}
