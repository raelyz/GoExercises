package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Println("key in a number")
	fmt.Scanln(&number)

	if number%2 != 0 {
		fmt.Println("odd number")
	} else {
		fmt.Println("even")
	}

	if number/10 >= 1 {
		fmt.Println("more than 1 digit")
	} else {
		fmt.Println("1 digit bruh")
	}

}
