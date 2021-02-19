package main

import (
	"fmt"
)

func main() {
	// for i := 0; i <= 1000; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 0; i <= 1000; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	for i := 0; i <= 1000; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
