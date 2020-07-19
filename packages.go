package main

import (
	"fmt"
	"math/rand"
	"time"
)

func packages() {
	rand.Seed(time.Now().Unix())
	fmt.Println("My favorite number is", rand.Intn(10))
}
