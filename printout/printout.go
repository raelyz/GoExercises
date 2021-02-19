package main

import (
	"fmt"
)

var (
	text           string  = "the following is the account information"
	firstName      string  = "Luke"
	lastName       string  = "Skywalkter"
	age            int64   = 20
	weight         float32 = 73.0
	height         float32 = 1.72
	credits        float32 = 123.55
	accName        string  = "admin"
	accountPW      string  = "password"
	subscribedUser bool    = true
)

func main() {

	fmt.Printf("%v - %T\n", text, text)
	fmt.Printf("%v - %T\n", firstName, firstName)
	fmt.Printf("%v - %T\n", lastName, lastName)
	fmt.Printf("%v - %T\n", age, age)
	fmt.Printf("%v - %T\n", weight, weight)
	fmt.Printf("%v - %T\n", height, height)
	fmt.Printf("%v - %T\n", credits, credits)
	fmt.Printf("%v - %T\n", accName, accName)
	fmt.Printf("%v - %T\n", accountPW, accountPW)
	fmt.Printf("%v - %T\n", subscribedUser, subscribedUser)
}
