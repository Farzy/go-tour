package main

import (
	"fmt"
	"go-tour/methods"
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
	title("Miscellaneous")
	misc()
	title("Types")
	moreTypes()
	title("Exercise: Slices")
	exerciceSlices()
	title("Maps")
	maps()
	title("Exercice: Maps")
	exerciceMaps()
	title("Function values and closures")
	functionValuesClosures()
	title("Exercice: Fibonacci closure")
	fibonacciClosure()
	title("Methods")
	methods.Main()
}
