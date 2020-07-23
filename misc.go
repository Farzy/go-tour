package main

import "fmt"

func misc() {
	const (
		pC     = 1 << iota // a control character.
		pP                 // a punctuation character.
		pN                 // a numeral.
		pS                 // a symbolic character.
		pZ                 // a spacing character.
		pLu                // an upper-case letter.
		pLl                // a lower-case letter.
		pp                 // a printable character according to Go's definition.
		pg     = pp | pZ   // a graphical character according to the Unicode definition.
		pLo    = pLl | pLu // a letter that is neither upper nor lower case.
		pLmask = pLo
	)

	fmt.Printf("pC = %v\n", pC)
	fmt.Printf("pP = %v\n", pP)
	fmt.Printf("pN = %v\n", pN)
	fmt.Printf("pS = %v\n", pS)
	fmt.Printf("pZ = %v\n", pZ)
	fmt.Printf("pLu = %v\n", pLu)
	fmt.Printf("pLl = %v\n", pLl)
	fmt.Printf("pp = %v\n", pp)
	fmt.Printf("pg = %v\n", pg)
	fmt.Printf("pLo = %v\n", pLo)
	fmt.Printf("pLmask = %v\n", pLmask)

}
