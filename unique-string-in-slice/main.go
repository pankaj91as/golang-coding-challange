package main

import "fmt"

func main() {
	stringSlice := []string{"pan", "kaj", "pan", "paw", "ar"}
	fmt.Print(removeDuplicateString(stringSlice))
}

func removeDuplicateString(str []string) []string {
	newSlice := []string{str[0]}
	for _, proStr := range str {
		exist := false
		for i := 0; i < len(newSlice); i++ {
			if newSlice[i] == proStr {
				exist = true
			}
		}
		if !exist {
			newSlice = append(newSlice, proStr)
		}
	}

	return newSlice
}
