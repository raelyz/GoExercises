package main

import "fmt"

func (merch game) printPrice() {
	fmt.Println(merch.price)
}
func (merch computerAccessories) printPrice() {
	fmt.Println(merch.price)
}
func (merch book) printPrice() {
	fmt.Println(merch.price)
}

type merchandise interface {
	printPrice()
}
