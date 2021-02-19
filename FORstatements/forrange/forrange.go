package main

import (
	"fmt"
)

func main() {
	var num1 int
	var num2 int
	fmt.Println("input first number")
	fmt.Scanln(&num1)
	fmt.Println("input second number")
	fmt.Scanln(&num2)

	for i := num1; i <= num2; i++ {
		if i%2 == 0 {
			fmt.Println(i, " even number")
		} else {
			fmt.Println(i, " odd number")
		}
	}

}
