package concurrency

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Try compiling and running the code with "safe = false", the runtime will complain about
// concurrent map writes:
// fatal error: concurrent map writes
//
// goroutine 14 [running]:
// runtime.throw(0x1127e90, 0x15)
//         /usr/local/Cellar/go/1.14.6/libexec/src/runtime/panic.go:1116 +0x72 fp=0xc000042728 sp=0xc0000426f8 pc=0x102fc12
// runtime.mapassign_faststr(0x11012c0, 0xc0000682a0, 0x112527e, 0x7, 0x0)
//         /usr/local/Cellar/go/1.14.6/libexec/src/runtime/map_faststr.go:211 +0x3f7 fp=0xc000042790 sp=0xc000042728 pc=0x10114d7
// go-tour/concurrency.(*SafeCounter).Inc(0xc000050380, 0x112527e, 0x7)
//         /Users/ffarid/src/GO/src/github.com/Farzy/go-tour/concurrency/sync-mutex.go:23 +0x4e fp=0xc0000427c8 sp=0xc000042790 pc=0x10cf04e
// runtime.goexit()
//         /usr/local/Cellar/go/1.14.6/libexec/src/runtime/asm_amd64.s:1373 +0x1 fp=0xc0000427d0 sp=0xc0000427c8 pc=0x105e3b1
// created by go-tour/concurrency.SyncMutex
//         /Users/ffarid/src/GO/src/github.com/Farzy/go-tour/concurrency/sync-mutex.go:44 +0x9c
//
// goroutine 1 [runnable]:
// time.Sleep(0x3b9aca00)
//         /usr/local/Cellar/go/1.14.6/libexec/src/runtime/time.go:174 +0x14e
// go-tour/concurrency.SyncMutex()
//         /Users/ffarid/src/GO/src/github.com/Farzy/go-tour/concurrency/sync-mutex.go:47 +0xbe
// main.main()
//         /Users/ffarid/src/GO/src/github.com/Farzy/go-tour/main.go:105 +0x3d8
var safe = true

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	if safe {
		c.mux.Lock()
	}
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	if safe {
		c.mux.Unlock()
	}
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	if safe {
		c.mux.Lock()
		// Lock so only one goroutine at a time can access the map c.v.
		defer c.mux.Unlock()
	}
	return c.v[key]
}

func SyncMutex() {
	if os.Getenv("SAFE") == "false" {
		safe = false
	}
	fmt.Println("Try setting 'SAFE=false' before running this function.")
	fmt.Println("Current safe value:", safe)

	c := SafeCounter{
		v: make(map[string]int),
	}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
