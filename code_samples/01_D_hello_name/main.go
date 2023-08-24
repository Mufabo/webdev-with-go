package main

import (
	"net/http"
	"fmt"
)

func getHelloHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Query().Get("name"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", getHelloHandleFunc)
	http.ListenAndServe(":8080", mux)
}