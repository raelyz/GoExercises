package main

import "fmt"

func main() {

}

type invalidSide struct {
	erroneousSide float64
	errorMessage  string
}

func (invalidSide *invalidSide) Error() string {
	format := "%v given for one of the sides. values must be more than 0."

	return fmt.Sprintf(format, invalidSide)
	return fmt.Sprintf("%v given for one of the sides. values must be more than 0.", invalidSide)
}
