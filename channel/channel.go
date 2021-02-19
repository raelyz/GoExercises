package main

import (
	"fmt"
	"runtime"
	"sync"
)

type bankBalance struct {
	amount   float64
	currency string
}

var (
	wg      sync.WaitGroup
	mutex   sync.Mutex
	account = &bankBalance{100, "SGD"}
)

func (b *bankBalance) deposit() {
	defer wg.Done()

	// mutex.Lock()
	// {
	var amount float64
	fmt.Print("Deposit Amount: ")
	fmt.Scanln(&amount)
	b.amount += amount
	fmt.Printf("Deposit Successfully: %v %v \n", b.currency, amount)
	b.display()
	// }
	// mutex.Unlock()
}

func (b *bankBalance) withdraw() {
	defer wg.Done()

	// mutex.Lock()
	// {
	var amount float64
	fmt.Print("Withdraw Amount: ")
	fmt.Scanln(&amount)
	if amount < b.amount {
		b.amount -= amount
		fmt.Println("Withdraw Successful")
		b.display()
	} else {
		fmt.Println("Withdraw Unsuccessful")
		b.display()
	}
	// }
	// mutex.Unlock()

}

func (b *bankBalance) display() {
	fmt.Printf("Balance: %v %2.f\n", b.currency, b.amount)
}

func (b *bankBalance) processBankBalance(choice int) {
	switch choice {
	case 1:
		b.deposit()
	case 2:
		b.withdraw()
	}
}

func main() {

	runtime.GOMAXPROCS(1)
	account.display()

	wg.Add(2)

	go account.processBankBalance(1)
	go account.processBankBalance(2)
	wg.Wait()
	account.display()
}
