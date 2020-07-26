package main

import (
	"fmt"
	"go-tour/errors"
	"go-tour/methods"
	more_types "go-tour/more-types"
	"go-tour/readers"
	"go-tour/utils"
	"os"
	"strings"
)

type Exercice struct {
	title    string
	function func()
}

var exercises = []Exercice{
	{"Hello", hello},
	{"Sandbox", sandbox},
	{"Packages", packages},
	{"Functions", functions},
	{"Variables", variables},
	{"Tyoes", types},
	{"Constants", constants},
	{"For loops", forLoop},
	{"If", ifs},
	{"Exercice: Loops and functions", exerciceLoopsAndFunctions},
	{"Switch", switches},
	{"Defer", defers},
	{"Miscellaneous", misc},
	{"More Types", moreTypes},
	{"Exercice: Slices", exerciceSlices},
	{"Maps", maps},
	{"Exercice: Maps", exerciceMaps},
	{"Function values and closures", functionValuesClosures},
	{"Exercice: Fibonacci closure", fibonacciClosure},
	{"Methods", methods.Main},
	{"Types manipulation", more_types.Main},
	{"Exercice: Stringer", exerciceStringer},
	{"Errors", errors.Main},
	{"Exercice: Errors", exerciceErrors},
	{"Readers", readers.Main},
	{"Exercice: Readers", exerciceReaders},
	{"Exercice: rot13Reader", exerciceRot13Reader},
}

func usage() {
	fmt.Printf(`Usage: %v [options] [args]

Options:
  --help | -h: This help
  --list | -l: List functions

Args:
  No argument: execute all functions
  func-name: execute only "func-name"
`, os.Args[0])
}

func funcName(s string) string {
	return strings.ToLower(
		strings.Replace(
			strings.Replace(s, " ", "-", -1),
			":", "", -1),
	)
}

func main() {
	if len(os.Args) == 1 {
		for _, ex := range exercises {
			utils.Title(ex.title)
			ex.function()
		}
	} else {
		arg := os.Args[1]
		if arg == "-h" || arg == "--help" {
			usage()
			return
		} else if arg == "-l" || arg == "--list" {
			println("Function list:")
			for _, ex := range exercises {
				fmt.Printf("  - %s\n", funcName(ex.title))
			}
		} else {
			for _, ex := range exercises {
				if arg == funcName(ex.title) {
					utils.Title(ex.title)
					ex.function()
					return
				}
			}
			fmt.Printf("Error: function '%v' not found.\n", arg)
		}
	}
}
