package main

import "fmt"

func main() {
	var num int
	for {
		fmt.Println("enter a number")
		fmt.Scanln(&num)

		if num%2 == 0 {
			fmt.Println("even!")
		} else {
			fmt.Println("odd!")
		}
	}
}
