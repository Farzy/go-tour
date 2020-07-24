package utils

import (
	"fmt"
	"strings"
)

func Title(title string) {
	fmt.Println("")
	fmt.Println(strings.Repeat("#", len(title)+8))
	fmt.Println("###", title, "###")
	fmt.Println(strings.Repeat("#", len(title)+8))
}

func Subtitle(subtitle string) {
	fmt.Println("")
	fmt.Println(subtitle)
	fmt.Println(strings.Repeat("-", len(subtitle)))
}
