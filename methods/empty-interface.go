package methods

import "fmt"

func emptyInterface() {
	var i interface{}

	describeEmpty(i)

	i = 42
	describeEmpty(i)

	i = "hello"
	describeEmpty(i)

}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
