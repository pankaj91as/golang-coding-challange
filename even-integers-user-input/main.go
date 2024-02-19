package main

import (
	"fmt"
)

func main() {
	var intSlice []int

	// Prompt the user to enter numbers
	fmt.Println("Enter integers separated by spaces (e.g., 1 2 3):")

	// Read input from the user until a non-integer value is entered
	var num int
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			break // Exit loop if a non-integer value is entered
		}
		intSlice = append(intSlice, num)
	}

	evenSlice := returnEvenIntOnly(intSlice)
	fmt.Println("Even numbers:", evenSlice)
}

func returnEvenIntOnly(intSlice []int) []int {
	newIntSlice := []int{}
	for _, v := range intSlice {
		if v%2 == 0 {
			newIntSlice = append(newIntSlice, v)
		}
	}
	return newIntSlice
}
