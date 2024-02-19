package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	index := searchInsert(nums, target)
	fmt.Printf("Value can insert @ position %d\n", index)
}

func searchInsert(nums []int, target int) int {
	index := 0
	for k, v := range nums {
		if v >= target {
			index = k
			break
		} else {
			index++
		}
	}
	return index
}
