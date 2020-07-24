package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount count the number of words in a string
func WordCount(s string) map[string]int {
	wcMap := make(map[string]int)
	for _, word := range strings.Fields(s) {
		wcMap[word]++
	}
	return wcMap
}

func exerciceMaps() {
	wc.Test(WordCount)
}
