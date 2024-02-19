package main

import "fmt"

func main() {
	str := "pankajshashikantpawar"
	numOfVow := countVowels(str)
	fmt.Printf("%d number Of vowels in provided string\n", numOfVow)
}

func countVowels(str string) int {
	count := 0
	for _, v := range str {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			count++
		}
	}
	return count
}
