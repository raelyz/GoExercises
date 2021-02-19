package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func hello(ch chan int) {
	time.Sleep(4 * time.Second)
	fmt.Println("Hello world goroutine", <-ch)
}
func main() {
	ch := make(chan int)
	for i := 1; i <= 4; i++ {

		go hello(ch)
		ch <- i
	}

	fmt.Println("main function")
}
