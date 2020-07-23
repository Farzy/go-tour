package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

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
