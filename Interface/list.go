package main

import "fmt"

func list(list []merchandise) {
	for _, merchandise := range list {
		fmt.Println(merchandise)
		merchandise.printPrice()
	}

}
