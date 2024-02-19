package main

import "fmt"

func main() {
	numbers := []int{36, 100, 9, 45, 67, 56, 3, 78}
	minmax := [2]int{numbers[0], numbers[0]}
	for _, v := range numbers {
		if minmax[0] > v {
			minmax[0] = v
		}
		if minmax[1] < v {
			minmax[1] = v
		}
	}
	fmt.Printf("Min: %d, Max: %d\n", minmax[0], minmax[1])
}
