package main

import (
	"fmt"
	"net/http"
	"strings"
	"strconv"
	//"fmt"
)

func getHelloHandleFunc(w http.ResponseWriter, r *http.Request) {
	//var bys []byte;
	// for k, v := range r.URL.Query(){
	// 	bys = append(bys, k...)
	// 	for _, val := range v{
	// 		bys = append(bys, val...)
	// 	}
	// }
	var sb strings.Builder
	for k, v := range r.URL.Query(){
		fmt.Fprintf(&sb, "%s:%s\n", k, v[0])
	}
	//w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(sb.String()))
}

func sumHandleFunc(w http.ResponseWriter, r *http.Request){
	sum := 0
	for _, v := range r.URL.Query(){
		val, _ := strconv.Atoi(v[0])
		sum += val
	}
	w.Write([]byte(strconv.Itoa(sum)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/params", getHelloHandleFunc)
	mux.HandleFunc("/sum", sumHandleFunc)
	mux.HandleFunc("/sum/", func(w http.ResponseWriter, r *http.Request){w.Write([]byte("sum"))})
	mux.HandleFunc("/sum/hello", func(w http.ResponseWriter, r *http.Request){w.Write([]byte("Hello"))})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){w.Write([]byte("Hello"))})
	http.ListenAndServe(":8080", mux)
}
