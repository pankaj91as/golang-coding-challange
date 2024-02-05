package main

import "fmt"

func main() {
	var searchNumber int = 4
	var foundCount = 0
	var i int = 0

	var providedArray = [5]int{10, 20, 30, 40, 50}

	for i = 0; i < len(providedArray); i++ {
		if providedArray[i] == searchNumber {
			fmt.Printf("Search Number %d found @ position %d\n", searchNumber, i+1)
			foundCount++
		}
	}

	if foundCount > 0 {
		fmt.Printf("Search Number %d found %d time\n", searchNumber, foundCount)
	} else {
		fmt.Printf("Provided number is not present in the array!\n")
	}
}
