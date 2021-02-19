package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Println("Key in a year")
	fmt.Scanln(&input)

	if input%4 == 0 {
		if input%400 == 0 || input%100 != 0 {
			fmt.Println("leap year!")
		} else {
			fmt.Println("NOT A LEAP YEAR")
		}

	} else {
		fmt.Println("NOT A LEAP YEAR")
	}

}
