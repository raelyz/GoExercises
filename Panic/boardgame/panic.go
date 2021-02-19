package main

import "fmt"

func main() {
	s := []string{"Carcassone", "Wildlife Safari", "Civilization"}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("You have entered an invalid choice. Value should be betweem 0 and 3")
		}
	}()
	for index, game := range s {
		fmt.Println(index+1, ":", game)
	}
	fmt.Println("What is your favourite game?")
	var input int
	fmt.Scanln(&input)
	fmt.Println("Oh I see. So your favourite game is :", s[input-1])
}
