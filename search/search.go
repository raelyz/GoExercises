package main

import "fmt"

func main() {
	data := []int{5, 6, 1, 2, 3, 15, 16, 7, 12, 11, 8}
	div(data[0:])
	fmt.Println(data)

}

func seqRec(data []int, target int, index int) {
	if index >= len(data) {
		fmt.Println("not found!")
		return
	}
	if data[index] != target {
		seqRec(data, target, index+1)
	} else {
		fmt.Println("found!", data[index], index)
	}

}

func binSer(data []int, start int, end int, target int) {
	mid := (start + end) / 2
	if data[mid] == target {
		fmt.Println("found!", data[mid])
	} else if data[mid] > target {
		binSer(data, start, mid-1, target)

	} else {
		binSer(data, mid+1, end, target)
	}

}

func insSort(data []int) {
	for i := 1; i < len(data); i++ {
		dat := data[i]
		l := i
		for l > 0 && data[l-1] > dat {
			data[l] = data[l-1]
			l--
		}
		data[l] = dat
	}
}

func selSort(data []int) {

	end := len(data)

	for i := 0; i < end; i++ {
		smallest := data[i]
		index := i
		fmt.Println(data)
		for k := i + 1; k < end; k++ {
			if data[k] < smallest {
				smallest = data[k]
				index = k
			}
			if k == end-1 {
				data[i], data[index] = data[index], data[i]

			}
		}

	}
}

func divide(data []int) []int {
	if len(data) == 1 {
		return data
	}
	mid := (len(data)) / 2
	left := divide(data[0:mid])
	right := divide(data[mid:len(data)])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	fmt.Println(left, right)
	size := len(left) + len(right)
	s := make([]int, size)
	var i, j int
	for index := range s {
		if i < len(left) && j < len(right) {
			if left[i] < right[j] {
				s[index] = left[i]
				i++
			} else {
				s[index] = right[j]
				j++
			}
		} else if i < len(left) {
			s[index] = left[i]
			i++
		} else {
			s[index] = right[j]
			j++
		}
	}

	fmt.Println(s)
	return s
}

func div(data []int) {
	if len(data) < 2 {
		return
	}

	mid := (len(data)) / 2
	right := len(data) - 1
	left := 0
	for left < right {
		fmt.Println(data)
		fmt.Println(left, right, mid)
		if data[left] < data[mid] {
			left++
		}
		if data[right] > data[mid] {
			right--
		}
		if data[left] > data[right] {
			data[left], data[right] = data[right], data[left]
		}
	}
	div(data[0:mid])
	div(data[mid:len(data)])
}
