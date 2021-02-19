package main

import (
	"fmt"
)

func main() {
	var firstNo int
	var secNo int

	fmt.Println("input first number")
	fmt.Scanln(&firstNo)
	fmt.Println("input second number")
	fmt.Scanln(&secNo)

	if firstNo < secNo {
		fmt.Printf("second number is bigger %v by %v", secNo, secNo-firstNo)

	} else {
		fmt.Printf("first number is bigger %v by %v", firstNo, firstNo-secNo)
	}
}
