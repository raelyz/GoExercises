package main

import (
	"fmt"
)

func main() {
	var unit int64
	var temp float64

	fmt.Println("1 for kelvin, 2 for celsius and 3 for fahrenheit")
	fmt.Scanln(&unit)
	fmt.Println("input current temperature")
	fmt.Scanln(&temp)
	if unit == 1 {
		temp1 := temp*(9.0/5) - 495.67
		temp2 := ((temp*(9.0/5) - 495.67) - 32)
		fmt.Println(temp1)
		fmt.Println(temp2)

	} else if unit == 2 {

		temp1 := temp*9.0/5 + 32
		temp2 := 5.0 / 9 * (temp*9/5 + 32 + 459.67)
		fmt.Println(temp1)
		fmt.Println(temp2)
	} else {
		temp1 := (temp - 32) * (5.0 / 9)
		temp2 := (5.0 / 9) * (temp + 459.67)
		fmt.Println(temp1)
		fmt.Println(temp2)
	}

}
