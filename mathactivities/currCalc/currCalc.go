package main

import (
	"fmt"
	"math"
)

func main() {

	var input1 float64
	var input05 float64
	var input02 float64
	var input01 float64
	var input005 float64

	fmt.Println("enter a 1 dollar coins amount:")
	fmt.Scanln(&input1)
	fmt.Println("enter a 50 cent coins amount:")
	fmt.Scanln(&input05)
	fmt.Println("enter a 20 cent coins amount:")
	fmt.Scanln(&input02)
	fmt.Println("enter a 10 cent coins amount:")
	fmt.Scanln(&input01)
	fmt.Println("enter a 5 cent coins amount:")
	fmt.Scanln(&input005)
	total := input1*1 + input05*0.5 + input02*0.2 + input01*0.1 + input005*0.05
	change := math.Mod(total, 2.0)
	fmt.Println(total)
	fmt.Println(change)
}
