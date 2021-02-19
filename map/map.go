package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := [][]string{{"USD", "US dollar", "1.1318"}, {"JPY", "Japanese yen", "121.05"}, {"GBP", "Pound Sterling", "0.90630"}, {"CNY", "Chinese yuan renminbi", "7.9944"}, {"SGD", "Singapore dollar", "1.5743"}, {"MYR", "Malaysian Ringgit", "4.8390"}}
	converter := make(map[string]currency)
	for _, curr := range slice {
		ex, _ := strconv.ParseFloat(curr[2], 64)
		details := currency{
			currency: curr[0],
			name:     curr[1],
			exr:      ex,
		}
		converter[curr[0]] = details
	}
	var firstCurr string
	var amount string
	var secCurr string

	fmt.Println("Key in first currency")
	fmt.Scanln(&firstCurr)
	fmt.Println("Key in amount")
	fmt.Scanln(&amount)
	fmt.Println("Key in second currency")
	fmt.Scanln(&secCurr)

	total, _ := strconv.ParseFloat(amount, 64)
	fmt.Println(total / converter[firstCurr].exr * converter[secCurr].exr)
}

type currency struct {
	currency string
	name     string
	exr      float64
}
