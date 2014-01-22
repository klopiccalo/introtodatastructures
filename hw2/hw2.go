package main

import (
	"fmt"

	"github.com/etherealmachine/introtodatastructures/hw2/magic"
)

var urls = []string{
	"http://en.wikipedia.org/wiki/Computer_science",
	"http://en.wikipedia.org/wiki/Computer_programming",
	"http://en.wikipedia.org/wiki/Measuring_programming_language_popularity",
}

func asOccuranceMap(words []string) map[string]int {
	occurences := make(map[string]int)
	return occurences
}

func reverseOccuranceMap(occurences map[string]int) (int, map[int][]string) {
	byCount := make(map[int][]string)
	max := 0
	return max, byCount
}

func main() {
	words := magic.GetWords(urls)

	// Create occurence map.
	occurences := asOccuranceMap(words)

	// Reverse occurence map.
	max, occurancesByCount := reverseOccuranceMap(occurences)

	// Print out occurences by count; largest first.
	for i := max; i > 1; i-- {
		if len(occurancesByCount[i]) > 0 {
			fmt.Printf("%d:\n", i)
			fmt.Println(occurancesByCount[i])
		}
	}
}
