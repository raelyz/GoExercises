package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	if area, err := calCircleArea(-1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area, "is the area")
	}

	if area, err := createTriangle(20, 5, 12); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area, "is the traingle area")
	}
	if area, err := createTriangle(20, 5, 20); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area, "is the traingle area")
	}
	if area, err := createTriangle(-20, 5, 20); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area, "is the traingle area")
	}
}

func calCircleArea(radius int) (int, error) {
	if radius < 0 {
		return 0, errors.New("Error Radius must be more than 0")
	}

	return int(math.Pi * math.Pow(float64(radius), 2)), nil

}

var invalidTriangleError = errors.New("Invalid Triangle")
var invalidSideError = errors.New("invalid side error")

func createTriangle(a int, b int, c int) (int, error) {
	if a < 0 || b < 0 || c < 0 {
		return -1, invalidSideError
	}
	if !(a+b > c && b+c > a && c+a > b) {
		return -1, invalidTriangleError
	}
	s := (a + b + c) / 2

	return int(math.Pow(float64(s*(s-a)*(s-b)*(s-c)), 0.5)), nil
}
