package methods

import "fmt"

type J interface {
	N()
}

func (t *T) N() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describeJ(j J) {
	fmt.Printf("(%v, %T)\n", j, j)
}

func nils() {
	var j J

	var t *T
	j = t
	describeJ(j)
	j.N()

	j = &T{"Hello"}
	describeJ(j)
	j.N()
}
