package concurrency

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		c <- x
	}
	close(c)
}

func RangeAndClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for n := range c {
		fmt.Println(n)
	}
}
