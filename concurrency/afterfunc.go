package concurrency

import (
	"fmt"
	"time"
)

func delayedFunc() {
	fmt.Printf("delayedFunc: This delayed function is called at %v\n", time.Now())
}
func UseAfterFunc() {
	fmt.Printf("Main: About to create a delayed func call at %v\n", time.Now())
	_ = time.AfterFunc(2*time.Second, delayedFunc)
	fmt.Println("Main: Sleeping for 3 seconds…")
	time.Sleep(3 * time.Second)
	fmt.Printf("Main: End sleep at %v\n", time.Now())
}
