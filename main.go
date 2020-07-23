package main

import (
	"fmt"
	"strings"
)

func title(title string) {
	fmt.Println("")
	fmt.Println(strings.Repeat("#", len(title)+8))
	fmt.Println("###", title, "###")
	fmt.Println(strings.Repeat("#", len(title)+8))
}

func subtitle(subtitle string) {
	fmt.Println("")
	fmt.Println(subtitle)
	fmt.Println(strings.Repeat("-", len(subtitle)))
}

func main() {
	title("Hello")
	hello()
	title("Sandbox")
	sandbox()
	title("Packages")
	packages()
	title("Functions")
	functions()
	title("Variables")
	variables()
	title("Types")
	types()
	title("Constants")
	constants()
	title("For loops")
	forLoop()
	title("If")
	ifs()
	title("Exercices: Loops and functions")
	exerciceLoopsAndFunctions()
	title("Switch")
	switches()
	title("Defer")
	defers()
	title("Types")
	moreTypes()
	title("Miscellaneous")
	misc()
}
