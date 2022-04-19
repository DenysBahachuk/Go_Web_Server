package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var count int32

func handleAll(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt32(&count, 1)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	temp := atomic.LoadInt32(&count)
	fmt.Println("Count:", temp)
	fmt.Fprintf(w, "<h1 align=\"center\">%d</h1>", count)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8080"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using default port number: ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}

	http.HandleFunc("/getCounter", getCounter)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", handleAll)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
