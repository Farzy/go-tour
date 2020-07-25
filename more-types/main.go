package more_types

import (
	"fmt"
	"go-tour/utils"
)

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
}
