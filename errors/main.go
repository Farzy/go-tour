package errors

import (
	"fmt"
	"go-tour/utils"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func Main() {
	utils.Subtitle("MyError")
	if err := run(); err != nil {
		fmt.Println(err)
	}

	utils.Subtitle("Error test on Atoi")
	s := "42Z"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Couldn't convert number %v: '%v'\n", s, err)
	} else {
		fmt.Println("Converted integer:", i)
	}
}
