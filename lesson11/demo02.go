package main

import (
	"fmt"
	"time"
)

// time
func main() {

}

// 获取当前时间 now
func time1() {
	// 返回值为Time结构体 ： 常量：日月年时分秒 周日-周六 方法：获取常量，计算。
	now := time.Now()

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	// 2023-2-23 20:40:31
	// Printf ： 整数补位--02如果不足两位，左侧用0补齐输出
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}

func time2() {
	//time1()
	// 打印时间
	now := time.Now()
	// 时间格式化 2023-02-23 20:43:49
	// 格式化模板： yyyy-MM-dd HH:mm:ss
	// Go语言诞生的时间作为格式化模板：2006年1月2号下午3点4分
	// Go语言格式化时间的代码：2006-01-02 15:04:05  （记忆方式：2006 12 345）
	// 固定的："2006-01-02 15:04:05"
	fmt.Println(now.Format("2006-01-02 15:04:05"))    // 24小时制
	fmt.Println(now.Format("2006-01-02 03:04:05 PM")) // 12小时制
	fmt.Println(now.Format("2006/01/02 15:04"))       // 2023/02/23 20:52
	fmt.Println(now.Format("15:04 2006/01/02"))       // 20:52 2023/02/23
	fmt.Println(now.Format("2006/01/02"))             // 2023/02/23
}

// 将字符串格式化为Time对象 (获取到网页传递的时间字符串，需要转化为Time才能在代码中使用)
func time3() {
	// 其他地方的时区格式：https://www.zeitverschiebung.net/cn/all-time-zones.html
	// 获取时间的时区 //  "Asia/Shanghai" 必须要大写 手动构建 ，如果不对，会报未知的时间错误
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 将字符串解析为时间 Time
	timeStr := "2023-02-23 20:53:08"
	// layout 格式 时间字符串  时区位置 , 需要和前端传递的格式进行匹配
	// func ParseInLocation(layout, value string, loc *Location)
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	fmt.Println(timeObj)
}

// 时间戳：更多时候和随机数结合
func time4() {
	// 格林威治时间自1970年1月1日(00:00:00 GMT)至当前时间的总秒数
	// 时间戳 Unix ： 1970.1.1 00:00:00 - 当下的一个毫秒数，Unix 时间戳，不会重复的。
	now := time.Now()
	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 纳秒的时间数

	fmt.Println(timestamp1)
	fmt.Println(timestamp2)
	//
	// 通过 Unix 转换time对象
	timeObj := time.Unix(timestamp1, 0) // 返回time对象
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	// 2023-2-23 20:40:31
	// Printf ： 整数补位--02如果不足两位，左侧用0补齐输出
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
