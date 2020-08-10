package interfaces

import (
	"fmt"
	"log"
	"net/http"
)

//type Handler interface {
//	ServerHTTP(w http.ResponseWriter, r *http.Request)
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

func ServeCounter() {
	ctr := new(Counter)
	http.Handle("/counter", ctr)

	notification := make(Chan, 10)
	http.Handle("/notification", notification)
	go notificationReceiver(notification)

	fmt.Println("Connect to:")
	fmt.Println("* http://localhost:8080/counter")
	fmt.Println("* http://localhost:8080/notification")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
