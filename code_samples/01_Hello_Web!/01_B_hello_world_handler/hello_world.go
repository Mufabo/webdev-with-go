package main

import "net/http"

type myHandleFunc func(http.ResponseWriter, *http.Request)

func (g myHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request){
	g(w, r)
}

var getHome myHandleFunc = func (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {


	http.Handle("/", getHome)
	http.ListenAndServe(":8080", nil)
}
