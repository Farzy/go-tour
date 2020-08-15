package goroutines

import (
	"fmt"
	"time"
)

func longRunningThread(c chan int) {
	fmt.Println("THREAD: This is a long running thread…")
	time.Sleep(3 * time.Second)
	fmt.Println("THREAD: End of long running thread.")
	c <- 1
}

func Main() {
	c := make(chan int)

	fmt.Println("MAIN: This is main.")
	fmt.Println("MAIN: Lets start a thread…")
	go longRunningThread(c)
	fmt.Println("MAIN: Now do something for a while…")
	time.Sleep(1 * time.Second)
	fmt.Println("MAIN: done")
	fmt.Println("MAIN: Let's wait for the long running thread…")
	<-c
	fmt.Println("MAIN: All done!")
}
