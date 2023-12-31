package main

import "fmt"

// 冒泡：每次筛选出一个最大或者最小的数.
/*
index   0   1   2   3   4
value   12  99  79  48  55
*/
// 冒泡排序逻辑，两两比较，大的往后移或者前移。 大
// 第一轮 ： 12 79 48 55 99 // 5
// 第二轮 ： 12 48 55 79 99 // 4
// 第三轮 ： 12 48 55 79 99 // 3 //
// 第四轮 ： 12 48 55 79 99 //
// 第五轮 ： 12 48 55 79 99

// 代码实践
/*
	// 两个数判断，如果一个数大，则交换位置，大放到后面
	if arr[x] > arr[x+1] {
		arr[x], arr[x+1] = arr[x+1],arr[x]
	}
	// 多轮判断，for， 循环次数 【数组大小】
*/
func main() {

	arr := [...]int{12, 99, 79, 48, 55, 1, 110, 111, 23, 52, 354, 2, 3412, 3, 12, 31}
	fmt.Println("初始数组：", arr)
	// 冒泡排序
	// 1、多少轮
	for i := 1; i < len(arr); i++ {
		// 2、筛选出来最大数字以后，我们下次不需要将它再计算了
		for j := 0; j < len(arr)-i; j++ {
			// 比较 // 改变升降序只需要改变符号即可
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}
		// 作业：排序已经结束，如果有个机制可以判断已经结束了，不需要再往后了。
		// if xxx {
		//	 break
		// }
		fmt.Println("第", i, "交换：", arr)
	}

}
