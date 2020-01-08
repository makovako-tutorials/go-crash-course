package main

import ("fmt"
		"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Printf("starting server")
	http.ListenAndServe(":3000", nil)
}