package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	tindex := searchInsert(nums, target)
	fmt.Printf("Value can insert @ position %d\n", tindex)
}

// Fourth iteration
func searchInsert(nums []int, target int) int {
	sI := 0
	lI := len(nums) - 1

	for sI <= lI {
		mid := (sI + lI) / 2
		currentVal := nums[mid]
		switch {
		case target < currentVal:
			lI = mid - 1
		case target > currentVal:
			sI = mid + 1
		default:
			return mid
		}
	}

	return sI
}

// Third iteration
// func searchInsert(nums []int, target int) int {
// 	sI := 0
// 	lI := len(nums) - 1

// 	for sI <= lI {
// 		mid := (sI + lI) / 2
// 		if target < nums[mid] {
// 			lI = mid - 1
// 		} else if target > nums[mid] {
// 			sI = mid + 1
// 		} else {
// 			return mid
// 		}
// 	}

// 	return sI
// }

// First iteration
// func searchInsert(nums []int, target int) int {
// 	index := 0
// 	for k, v := range nums {
// 		if v >= target {
// 			index = k
// 			break
// 		} else {
// 			index++
// 		}
// 	}
// 	return index
// }

// Second iteration
// func searchInsert(nums []int, target int) int {
// 	pos := findPosition(nums, target, 0)
// 	return pos
// }

// func findPosition(nums []int, target int, index int) int {
// 	if index == len(nums) {
// 		return index
// 	}

//		if nums[index] >= target {
//			return index
//		} else {
//			index += 1
//			return findPosition(nums, target, index)
//		}
//	}
