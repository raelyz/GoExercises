package main

import (
	"math/rand"
	"time"
)

func random(num int) int {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(num)
	return x
}
