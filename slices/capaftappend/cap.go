package main

import "fmt"

func main() {

	slice := make([]float64, 4, 5)
	slice[0] = 9.50
	slice[1] = 8.00
	slice[2] = 9.8
	slice[3] = 7.40
	slice = append(slice, 8.40)
	slice = append(slice, 9.40)
	slice = append(slice, 7.20)

	fmt.Println(slice)
}
