package main

import "fmt"

func main() {
	intSlice := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 12}
	evenSlice := returnEvenIntOnly(intSlice)
	fmt.Print(evenSlice, "\n")
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
