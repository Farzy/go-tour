package allocation

import "fmt"

func Main() {

	const (
		Enone = iota + 1
		Eio
		Einval
	)

	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

	fmt.Printf("a (%T) = %[1]v\n", a)
	fmt.Printf("s (%T) = %[1]v\n", s)
	fmt.Printf("s (%T) = %[1]v\n", m)
}
