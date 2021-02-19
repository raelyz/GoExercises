package main

import (
	"fmt"
)

func main() {
	var num int
	for {
		fmt.Println("input a number")
		fmt.Scanln(&num)
		if num%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
		if num/10 != 0 {
			fmt.Println("more than 1 digit")
		} else {
			fmt.Println("1 digit")
		}

	}
}
