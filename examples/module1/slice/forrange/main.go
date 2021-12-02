package main

import (
	"fmt"
)

func main() {
	mySlice := []int{10, 20, 30, 40, 50}
	// go语言是值传递
	for _, value := range mySlice {
		value *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)

	for index := range mySlice {
		mySlice[index] *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)
}
