package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Println("input num1")
	fmt.Scanln(&num1)
	fmt.Println("input num2")
	fmt.Scanln(&num2)

	for i := num1; i <= num2; i++ {
		fmt.Println(i)
	}
	for i := num2; i >= num1; i-- {
		fmt.Println(i)
	}

}
