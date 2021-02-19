package main

import "fmt"

func main() {

	slice := []int{20, 21, 23, 25, 22, 27, 23, 25, 20, 30, 24, 22, 23, 24, 22}

	slice1 := slice[0:4]
	slice2 := slice[5:10]
	slice3 := slice[11:15]
	slice4 := [][]int{slice1, slice2, slice3}
	for _, room := range slice4 {
		var total int = 0
		var count int = 0
		for _, temp := range room {
			total += temp
			count++
		}
		fmt.Println(float64(total) / float64(count))
	}
}
