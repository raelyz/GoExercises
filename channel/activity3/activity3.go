package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	done := time.After(1 * time.Second)
	echo := make(chan []string)
	go write(echo, "1")
	go write(echo, "2")
	for {
		select {
		case buf := <-echo:
			fmt.Println(buf)
		case <-done:
			fmt.Println("Timed out")
			os.Exit(0)
		}
	}
}

func write(ch chan []string, str string) {
	file, err := os.Open("game" + str + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	slice := make([]string, 5)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}
	ch <- slice
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
