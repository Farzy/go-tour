package main

import "fmt"

func title(title string) {
	fmt.Println("")
	fmt.Println("###################")
	fmt.Println("###", title, "###")
	fmt.Println("###################")
}
func main() {
	title("Hello")
	hello()
	title("Sandbox")
	sandbox()
	title("Packages")
	packages()
	title("Functions")
	functions()
	title("Variables")
	variables()
	title("Types")
	types()
}
