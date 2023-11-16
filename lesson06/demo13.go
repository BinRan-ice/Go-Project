package main

import "fmt"

// map 结合 slice 来使用

/*
需求：
1、创建map来存储人的信息，name，age,sex，addr
2、每个map保存一个的信息
3、将这些map存入到切片中
4、打印这些数据
*/
func main() {

	user1 := make(map[string]string)
	user1["name"] = "binran"
	user1["age"] = "27"
	user1["sex"] = "男"
	user1["addr"] = "重庆"

	user2 := make(map[string]string)
	user2["name"] = "feifei"
	user2["age"] = "30"
	user2["sex"] = "男"
	user2["addr"] = "长沙"

	user3 := map[string]string{"name": "小蓝", "age": "18", "sex": "男", "addr": "火星"}
	fmt.Println(user3)

	// 3个数据有了，存放到切片中，供我们使用
	userDatas := make([]map[string]string, 0, 3)
	userDatas = append(userDatas, user1)
	userDatas = append(userDatas, user2)
	userDatas = append(userDatas, user3)

	fmt.Println(userDatas)

	// 0 map[string]string
	for _, user := range userDatas {
		//fmt.Println(i)
		if user["name"] == "morenbu" {
			fmt.Println(user["addr"])
		}
	}

}
