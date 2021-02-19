package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var input string
	fmt.Println("input height")
	fmt.Scanln(&input)
	height, _ := strconv.Atoi(input)
	fmt.Println("input weight")
	fmt.Scanln(&input)
	weight, _ := strconv.Atoi(input)

	switch bmi := float64(weight) / math.Pow(float64(height)/100, 2); {
	case bmi < 18.5:
		fmt.Println("underweight")
	case bmi >= 18.5 && bmi <= 24.9:
		fmt.Println("healthy")
	case bmi >= 25.0 && bmi <= 29.9:

		fmt.Println("over")
	case bmi >= 30.0 && bmi <= 34.9:
		fmt.Println("obese")
	case bmi >= 35.0 && bmi <= 39.9:
		fmt.Println("sever obese")
	default:

		fmt.Println("morbid obese")
	}

}
