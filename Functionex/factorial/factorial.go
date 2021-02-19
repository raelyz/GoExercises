package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("enter a number")
	fmt.Scanln(&num)
	fmt.Println(factorial(num))
}

func factorial(num int) int {
	if num == 1 || num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
