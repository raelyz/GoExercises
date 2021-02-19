package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "this is a string"
	b := 12345
	c := 1.2345
	d := true
	inspect(a)
	inspect(b)
	inspect(c)
	inspect(d)
}

func inspect(input input) {
	fmt.Println(reflect.TypeOf(input))
	fmt.Println(reflect.ValueOf(input))
}

type input interface {
}
