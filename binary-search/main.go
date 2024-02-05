package main

import "fmt"

func main() {
	var providedArray = []int{10, 20, 30, 40, 50, 60, 70, 80}

	searchNumber, length, zero, one, i, mid, foundCount := 5, len(providedArray), 0, len(providedArray)-1, 0, 0, 0
	for i = 0; i < length; i++ {
		mid = (zero + one) / 2
		if searchNumber < providedArray[mid] {
			one = mid
		} else if searchNumber == providedArray[mid] {
			fmt.Printf("Provided number found @ position %d\n", mid)
			foundCount++
			break
		} else {
			zero = mid
		}
	}

	if foundCount == 0 {
		fmt.Println("Provided number not present in the array!")
	}
}
