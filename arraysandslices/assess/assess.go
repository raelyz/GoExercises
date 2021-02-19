package main

import "fmt"

func main() {
	var os = [6]string{"Windows XP", "Linux 1.0", "Raspbian Teensy", "MacOS", "IOS", "Google Android"}

	os[0] = "Windows 10"
	os[1] = "Linux 16.04"
	os[2] = "Raspbian Buster"

	for _, osName := range os {
		fmt.Println(osName)
	}
}
