package main

import "fmt"

func main() {

	slice := make([]float64, 4, 5)
	slice[0] = 9.50
	slice[1] = 8.00
	slice[2] = 10.20
	slice[3] = 7.40
	fmt.Println(slice[2])
	slice[2] = 9.80
	fmt.Println(slice)
}
