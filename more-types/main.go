package more_types

import (
	"fmt"
	"go-tour/utils"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func Main() {
	utils.Subtitle("Type assertions")

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	fmt.Println("Uncomment line in source code to trigger panic.")
	// panic: interface conversion: interface {} is string, not float64
	//f = i.(float64)
	//fmt.Println(f)

	utils.Subtitle("Type switches")
	do(21)
	do("hello")
	do(true)

	utils.Subtitle("Stringer")
	a := Person{
		"Arthur Dent",
		42,
	}
	z := Person{
		"Zaphod Beeblebrox",
		9001,
	}
	fmt.Println(a, z)
}
