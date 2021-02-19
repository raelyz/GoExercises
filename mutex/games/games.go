package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var (
	slice = make([]game, 10)
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go manageGameItems()
	go manageGameItems()
	wg.Wait()
}

type game struct {
	name  string
	price int
}

func add() {
	var name string
	var price int
	fmt.Println("game name")
	fmt.Scanln(&name)
	fmt.Println("price")
	fmt.Scanln(&price)
	newGame := game{
		name, price,
	}
	slice = append(slice, newGame)
}

func manageGameItems() {
	defer wg.Done()
	mutex.Lock()
	{
		var input string
		fmt.Println("1. display games")
		fmt.Println("2. add games")
		fmt.Scanln(&input)
		if input, err := strconv.Atoi(input); err != nil {
			fmt.Println(err)
		} else {
			if input == 1 {
				for games := range slice {
					fmt.Println(games)
				}
			} else if input == 2 {
				add()
			} else {
				fmt.Println("1 or 2 only")
			}
		}
	}
	mutex.Unlock()
}
