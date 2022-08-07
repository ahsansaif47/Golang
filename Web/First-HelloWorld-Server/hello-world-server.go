package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	// writer and the content you need to write to response writer(in Fprint arguments)..
	fmt.Fprint(w, "Hello World!")
}

func main() {
	// HandleFunc associates a path with a webpage function..
	http.HandleFunc("/", helloWorldPage)

	// starting the server by calling ListenAndServe..
	/*
		"" in ListenAndServe refers to port 80 and nil handler
		uses default handler to serve most  for routing
	*/
	http.ListenAndServe("", nil)

}
