package main

import (
	"go-tour/errors"
	"go-tour/methods"
	more_types "go-tour/more-types"
	"go-tour/readers"
	"go-tour/utils"
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
	{"Exercices: Loops and functions", exerciceLoopsAndFunctions},
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

func main() {
	for _, ex := range exercises {
		utils.Title(ex.title)
		ex.function()
	}
}
