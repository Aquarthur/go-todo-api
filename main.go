package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	hello := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hullo there!"))
	}
	mux.Handle("/hello", http.HandlerFunc(hello))
	http.ListenAndServe(":8080", mux)
}
