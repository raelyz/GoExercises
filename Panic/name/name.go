package main

import "fmt"

var (
	first string
	last  string
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Please enter your first name:")
	fmt.Scanln(&first)
	fmt.Println("Please enter your last name:")
	fmt.Scanln(&last)
	fmt.Println(first)
	if first == "" || last == "" {
		panic("last name cannot be empty")
	}
}
