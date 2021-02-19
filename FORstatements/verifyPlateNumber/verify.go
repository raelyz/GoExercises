package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var input string
	fmt.Println("input plate number")
	fmt.Scanln(&input)
	alphabetMap := make(map[byte]int)
	var alpha byte = 'A'
	for i := 1; i <= 26; i++ {
		alphabetMap[alpha] = i
		alpha++
	}

	index := indexOfLastAlphabet(input)

	alph := input[:index]
	nums := input[index:]

	alphLength := len(alph)
	var decode []int
	if alphLength == 1 {
		decode = append(decode, 0)
		decode = append(decode, alphabetMap[alph[0]])
	} else {
		for i := alphLength - 1; i > alphLength-3; i-- {
			decode = append(decode, alphabetMap[alph[i]])
		}
	}
	for i := 0; i < 4; i++ {
		decode = append(decode, 0)
	}
	var individualNums []string = strings.Split(nums, "")
	var g int = 5
	for i := len(nums) - 1; i >= 0; i-- {
		decode[g], _ = strconv.Atoi(individualNums[i])
		g--
	}
	var total int
	for _, s := range decode {
		fmt.Println(s)
		total += s
	}
	total *= 27
	total %= 19
	checkMap := make(map[int]byte)
	checkMap[0] = 'A'
	checkMap[1] = 'Z'
	checkMap[2] = 'Y'
	checkMap[3] = 'X'
	checkMap[4] = 'U'
	checkMap[5] = 'T'
	checkMap[6] = 'S'
	checkMap[7] = 'R'
	checkMap[8] = 'P'
	checkMap[9] = 'M'
	checkMap[10] = 'L'
	checkMap[11] = 'K'
	checkMap[12] = 'J'
	checkMap[13] = 'H'
	checkMap[14] = 'G'
	checkMap[15] = 'E'
	checkMap[16] = 'D'
	checkMap[17] = 'C'
	checkMap[18] = 'B'

	if checkMap[total] == alph[len(alph)-1] {
		fmt.Println("correct!")
	} else {
		fmt.Println("incorrect")
	}

}

func indexOfLastAlphabet(s string) int {
	var index int
	for i, r := range s {
		if !unicode.IsLetter(r) {
			index = i
			break
		}
	}
	return index
}
