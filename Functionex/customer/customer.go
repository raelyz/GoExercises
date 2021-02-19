package main

import "fmt"

func (person customer) userCredentials() {
	fmt.Println(person.passWord, person.userName)
}

func (person customer) userAddress() {
	fmt.Println(person.address)
}

func (person customer) printAll() {
	fmt.Println(person)
}

type customer struct {
	firstName string
	lastName  string
	userName  string
	passWord  string
	email     string
	phone     int
	address   string
	_         struct{}
}
