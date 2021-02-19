package main

import "fmt"

func list(list []game) {
	for _, games := range list {
		fmt.Println(games)
	}
}
