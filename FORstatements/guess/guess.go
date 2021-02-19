package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	var input int

	x := random(100)
	i := 0
	for i < 5 {
		fmt.Println("guess?")
		fmt.Scanln(&input)
		if input == x {
			fmt.Println("correct!")
			break
		} else if input == 101 {
			fmt.Println("thanks for playing!", text)
			break
		} else {
			fmt.Println("wrong! guess again")
			i++
		}
	}
}
