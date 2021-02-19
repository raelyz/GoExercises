package main

import (
	"fmt"
	"strconv"
)

func main() {
	list1 := [4]string{"ans", "wer", "is", "of"}
	list2 := [4]string{"-", "+", "*", "/"}
	list3 := [4]string{"5", "10", "4", "0"}
	result, err := strconv.Atoi(list3[0])
	result2, err := strconv.Atoi(list3[2])
	if err != nil {
		fmt.Println(err)
	}
	total := result + result2
	final := strconv.Itoa(total)
	fmt.Println(list1[0] + list1[1] + " " + list1[3] + " " + list3[0] + " " + list2[1] + " " + list3[2] + " " + list1[2] + " " + final)
}
