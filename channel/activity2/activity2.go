package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)
	a := requestInput()
	res := make(chan int, 1)
	ch := make(chan int, 1)
	var total int
	for i := 2; i <= 3; i++ {
		go pow(ch, res, i)
		ch <- a
		total += <-res
	}

	wg.Wait()
	fmt.Println(total)
	// fmt.Println(<-res, <-res)
}

func requestInput() int {
	var input string
	fmt.Println("Key in a num")
	fmt.Scanln(&input)
	if option, err := strconv.Atoi(input); err != nil {
		fmt.Println(err)
		return requestInput()
	} else {
		if option >= 100 && option <= 999 {
			return option
		}
		fmt.Println(errors.New("3 DIGITS ONLY"))
		return requestInput()
	}
}

func pow(ch chan int, res chan int, pow int) {
	defer wg.Done()
	a := <-ch
	var result float64
	for a != 0 {
		digit := a % 10
		a = a / 10
		result += math.Pow(float64(digit), float64(pow))
	}
	res <- int(result)
}
