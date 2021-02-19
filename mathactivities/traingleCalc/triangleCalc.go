package main

import (
	"fmt"
	"math"
)

func main() {
	var first float64
	fmt.Println("key in first length")
	fmt.Scanln(&first)
	var second float64
	fmt.Println("key in second length")
	fmt.Scanln(&second)
	var angle float64
	fmt.Println("key in angle")
	fmt.Scanln(&angle)

	result := math.Pow(math.Pow(first, 2)+math.Pow(second, 2)-2*first*second*math.Cos(angle), 0.5)

	fmt.Println(result)
}
