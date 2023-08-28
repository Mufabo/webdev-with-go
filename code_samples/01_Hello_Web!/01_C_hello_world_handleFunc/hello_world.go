package main

import "net/http"

type myHandler struct {}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}

func main() {

	http.Handle("/", &myHandler{})
	http.ListenAndServe(":8080", nil)
}