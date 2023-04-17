package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func indexx(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is index!!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf("Started %s %s", r.Method, r.URL.Path)
	fmt.Fprintf(w, "This is hello!!")

	log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
}

func ServerMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/", indexx)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
