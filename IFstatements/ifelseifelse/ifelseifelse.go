package main

import (
	"fmt"
)

var (
	accounts = [3]string{"Admin", "Robin", "John"}
)

func main() {
	var input string
	var found bool = false
	fmt.Println("Input account name")
	fmt.Scanln(&input)

	for _, s := range accounts {
		if input == s {
			found = true
		}
	}

	if found == true {
		fmt.Printf("Welcome %v", input)
	} else {
		fmt.Printf("input a valid user")
	}

}
