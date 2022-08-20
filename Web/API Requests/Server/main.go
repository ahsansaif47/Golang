package main

import (
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/url":
		rootListen(w, r)
	default:
		rootListen(w, r)
	}
}

func rootListen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe("", nil)
}
