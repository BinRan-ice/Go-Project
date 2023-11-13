package main

import "fmt"

/*
* * * * *
* * * * *
* * * * *
* * * * *
* * * * *
 */
func main() {
	// for循环的嵌套
	for i := 1; i <= 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
