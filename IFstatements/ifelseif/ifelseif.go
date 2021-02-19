package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Println("enter a number")
	fmt.Scanln(&number)

	if (number%5 == 0) && (number%6 == 0) {
		fmt.Printf("divisible \n")
	} else {
		fmt.Println("not divisble")
	}
}
