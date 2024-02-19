package main

import "fmt"

func main() {
	reverseSlice := []string{"Pawar", "Shashikant", "Pankaj"}
	newStringSlice := reverseProvidedSlice(reverseSlice)
	fmt.Println(newStringSlice)
}

func reverseProvidedSlice(rs []string) []string {
	newFS := []string{}
	for i := len(rs) - 1; i >= 0; i-- {
		newFS = append(newFS, rs[i])
	}
	return newFS
}
