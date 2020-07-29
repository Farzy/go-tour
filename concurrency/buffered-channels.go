package concurrency

import "fmt"

func BufferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// Uncomment to deadlock with "fatal error: all goroutines are asleep - deadlock!"
	// Buffer is full here, writing to the channel would block
	//ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// Uncomment to deadlock with "fatal error: all goroutines are asleep - deadlock!"
	// Buffer is empty here, reading from the channel would block
	//fmt.Println(<- ch)
}
