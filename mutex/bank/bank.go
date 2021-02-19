package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var (
	bank  = bankBalance{0, "SGD"}
	mutex sync.Mutex
	wg    sync.WaitGroup
	bank2 bankBalance
)

func main() {
	bank2.testing()
	fmt.Println(bank2)
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go processBankBalance()
	go processBankBalance()
	wg.Wait()
}

type bankBalance struct {
	amount   float64
	currency string
}

func deposit() {
	var input string
	fmt.Println("how much to deposit")
	fmt.Scanln(&input)
	if amount, err := strconv.ParseFloat(input, 64); err != nil {
		fmt.Println(err)
	} else {
		bank.amount += amount
	}
}
func (b *bankBalance) testing() {
	b.amount = 25
	fmt.Println(b.amount)
}

func withdraw() {
	var input string
	fmt.Println("how much to withdraw")
	fmt.Scanln(&input)
	if amount, err := strconv.ParseFloat(input, 64); err != nil {
		fmt.Println(err)
	} else {
		bank.amount -= amount
	}
}

func display() {
	fmt.Println(bank)
}

func processBankBalance() {
	defer wg.Done()
	mutex.Lock()
	{
		var input string
		fmt.Println("1 to depo")
		fmt.Println("2 to withdraw")
		fmt.Println("3 to display")
		fmt.Scanln(&input)
		if amount, err := strconv.Atoi(input); err != nil {
			fmt.Println(err)
		} else {
			switch amount {
			case 1:
				deposit()
			case 2:
				withdraw()
			case 3:
				display()
			}
		}
	}
	mutex.Unlock()
}
