package main

import "fmt"

func main() {

	var guess int
	var guess2 int
	var conditions bool = true
	var operator string
	for conditions {
		fmt.Println("Enter an integer value: ")
		fmt.Scanln(&guess)
		fmt.Println("Enter operator ")
		fmt.Scanln(&operator)
		fmt.Println("Enter a second integer value: ")
		fmt.Scanln(&guess2)
		if operator == "+" {
			fmt.Println(guess + guess2)
		} else if operator == "-" {
			fmt.Println(guess - guess2)
		} else if operator == "/" {
			fmt.Println(guess / guess2)
		} else if operator == "*" {
			fmt.Println("Result")
			fmt.Println(guess * guess2)
		}
	}
}
