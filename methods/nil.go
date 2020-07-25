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

func nilInterfaceValue() {
	var k I
	fmt.Println("Uncomment line in source code to trigger panic.")
	describe(k)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x10c5863]
	//k.M()
}
