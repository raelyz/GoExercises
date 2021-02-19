package main

import "fmt"

func main() {
	customer1 := customer{
		Firstname:       "Annakin",
		Lastname:        "Skywalker",
		age:             45,
		subscriber:      true,
		homeaddress:     "Death Star",
		phone:           1234567,
		creditavailable: 10000.00,
		currcartcost:    0.00,
		currordercost:   0.00,
	}
	customer2 := customer{
		Firstname:       "Han",
		Lastname:        "Solo",
		age:             50,
		subscriber:      false,
		homeaddress:     "Tatooine",
		phone:           4321765,
		creditavailable: 20000.00,
		currcartcost:    0.00,
		currordercost:   0.00,
	}
	fmt.Println(customer1.Firstname)
	fmt.Println(customer2)
}

type customer struct {
	Firstname       string
	Lastname        string
	age             int
	subscriber      bool
	homeaddress     string
	phone           int
	creditavailable float32
	currcartcost    float32
	currordercost   float32
}
