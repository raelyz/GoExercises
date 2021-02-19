package main

import "fmt"

func main() {
	var os = [24]float64{20.1, 24.0, 27.3, 30.1, 26.4, 22.2, 20.1, 24.0, 27.3, 30.1, 26.4, 20.1, 24.0, 27.3, 30.1, 26.4, 20.1, 24.0, 27.3, 30.1, 26.4, 20.1, 24.0, 27.3}
	var total float64
	for _, temp := range os {
		total += temp
	}
	var avg float64 = total / 24

	fmt.Println(avg)
}
