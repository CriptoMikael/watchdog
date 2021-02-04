package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var lastTime time.Time
var c chan bool

func ping(w http.ResponseWriter, req *http.Request) {
	c <- true
}

func main() {
	timeout := flag.Duration("timeout", 30*time.Second, "Timeout for ping in sec (default: 30 sec)")
	port := flag.String("port", "8090", "Port for check")
	flag.Parse()

	c = make(chan bool)

	http.HandleFunc("/", ping)
	go func() {
		http.ListenAndServe(":"+*port, nil)
	}()

	for {
		select {
		case <-c:
			continue
		case <-time.After(*timeout):
			fmt.Println("No answer 10 sec")
		}
	}
}
