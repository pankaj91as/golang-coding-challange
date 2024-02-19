package main

import (
	"fmt"
	"sort"
)

func sortTree(tree []int) []int {
	sort.Ints(tree)
	return tree
}

func main() {
	var givenTree = []int{10, 1, 2, 3, 4, 20, 30, 40, 50}
	fmt.Println(sortTree(givenTree))
}
