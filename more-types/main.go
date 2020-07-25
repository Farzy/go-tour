package more_types

import (
	"fmt"
	"go-tour/utils"
)

func Main() {
	utils.Subtitle("Type assertions")

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panic: interface conversion: interface {} is string, not float64
	//f = i.(float64)
	//fmt.Println(f)
}
