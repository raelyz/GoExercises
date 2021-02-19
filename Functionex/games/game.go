package main

import "fmt"

func (game game) printPrice() {
	fmt.Println(game.price)
}

type game struct {
	title string
	price float64
}
