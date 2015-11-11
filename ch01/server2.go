// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/count", counter)
	http.HandleFunc("/", handler) // each request calls handler
	// BobK:  This is just fucked, under what conditions do we log.Fatal here?
	// I can't help but feel I am already looking at bad 'Go'
	// Ha, ha.  Port 8000 is already in use by dynamodb-local, but no error.
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
