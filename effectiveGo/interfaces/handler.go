package interfaces

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//type Handler interface {
//	ServerHTTP(w http.ResponseWriter, r *http.Request)
//}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers.  If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler object that calls f.
//type HandlerFunc func(http.ResponseWriter, *http.Request)

// ServeHTTP calls f(w, req).
//func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	f(w, req)
//}

// Simple counter server
type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	ctr.n++
	_, _ = fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	_, _ = fmt.Fprintf(w, "notification sent")
}

func notificationReceiver(ch Chan) {
	for m := range ch {
		log.Printf("Request = %v\n", *m)
	}
}

// Argument server.
func ArgServer(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, os.Args)
}

func ServeCounter() {
	ctr := new(Counter)
	http.Handle("/counter", ctr)

	notification := make(Chan, 10)
	http.Handle("/notification", notification)
	go notificationReceiver(notification)

	http.HandleFunc("/args", http.HandlerFunc(ArgServer))

	fmt.Println("Connect to:")
	fmt.Println("* http://localhost:8080/counter")
	fmt.Println("* http://localhost:8080/notification")
	fmt.Println("* http://localhost:8080/args")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
