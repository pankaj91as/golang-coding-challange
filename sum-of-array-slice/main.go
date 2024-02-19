package main

import "fmt"

func main() {
	numbers := []int{34, 67, 23, 89, 65, 1, 90, 23, 45, 87}
	sum := addition(numbers)
	fmt.Println("SUM:", sum)
}

func addition(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
