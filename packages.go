package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func packages() {
	rand.Seed(time.Now().Unix())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
