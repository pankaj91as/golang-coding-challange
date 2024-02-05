package main

import "fmt"

func main() {
	var providedArray = []int{10, 20, 30, 40, 50, 60, 70, 80}

	searchNumber, zero, one, foundCount := 60, 0, len(providedArray)-1, 0

	// foundCount = binarySearchRecursive(providedArray, searchNumber, zero, one)
	foundCount = binarySearchIterative(providedArray, searchNumber, zero, one)

	if foundCount != -1 {
		fmt.Printf("Element %d found at index %d\n", searchNumber, foundCount)
	} else {
		fmt.Printf("Element %d not found in the array\n", searchNumber)
	}
}

// Iterative Implementation
func binarySearchIterative(arr []int, searchNumber int, zero int, one int) int {
	var i int
	var length = len(arr)
	var mid int
	for i = 0; i < length; i++ {
		mid = (zero + one) / 2
		if searchNumber < arr[mid] {
			one = mid
		} else if searchNumber == arr[mid] {
			break
		} else {
			zero = mid
		}
	}

	return mid
}

// Recursive Implementation
func binarySearchRecursive(arr []int, searchNumber int, zero int, one int) int {
	if zero > one {
		return -1
	}

	mid := (zero + one) / 2

	if arr[mid] < searchNumber {
		return binarySearchRecursive(arr, searchNumber, mid+1, one)
	} else if arr[mid] == searchNumber {
		return mid
	} else {
		return binarySearchRecursive(arr, searchNumber, zero, mid-1)
	}
}
