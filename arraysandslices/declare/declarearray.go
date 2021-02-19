package main

import "fmt"

func main() {
	var os = [6]string{"Windows XP", "Linux 1.0", "Raspbian Teensy", "MacOS", "IOS", "Google Android"}

	for _, osName := range os {
		fmt.Println(osName)
	}
}
