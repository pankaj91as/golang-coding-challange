package main

import "testing"

func BenchmarkSearchInsert(b *testing.B) {
	nums := []int{1, 3, 5, 6}
	target := 7
	searchInsert(nums, target)
}
