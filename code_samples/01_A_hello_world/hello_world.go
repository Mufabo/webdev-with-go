package main

import "net/http"

func getHomeHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {


	http.HandleFunc("/", getHomeHandleFunc)
	http.ListenAndServe(":8080", nil)
}
