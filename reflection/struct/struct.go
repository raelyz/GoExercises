package main

import (
	"fmt"
	"reflect"
)

func main() {
	customer1 := customer{
		fName:        "John",
		lName:        "Wick",
		userID:       123123123,
		invoiceTotal: 10000,
	}

	customer2 := reflect.ValueOf(&customer1).Elem()
	typeOfC := customer2.Type()
	for i := 0; i < customer2.NumField(); i++ {
		f := customer2.Field(i)
		fmt.Println(typeOfC.Field(i).Name, f.Type(), f.Interface())
	}
}

type customer struct {
	fName        string
	lName        string
	userID       int
	invoiceTotal float64
}
