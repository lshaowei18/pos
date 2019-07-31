package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Echo)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln("Can't start server")
	}
}

func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You asked to %s %s\n", r.Method, r.URL.Path)
}
