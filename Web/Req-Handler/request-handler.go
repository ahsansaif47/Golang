/*
	Objective..
	Defining path level works:
	What a path is supposed to show or do..
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleFunction(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Println("Printing plain Hello Ahsan!")
		fmt.Fprint(w, "Hello Ahsan!")
	case "/heading-hello":
		headerHello(w, r)
	case "/ninja":
		fmt.Fprint(w, "Hello Ninja!")
	case "/plain-content":
		setContentTypePlain(w, r)
	case "/timeout":
		timeoutFunc(w, r)
	case "/stealth-mode":
		HelloWorldNinja(w, r)
	default:
		fmt.Fprint(w, "Got an Error!")
	}
	// Checking the request type from http.Request instance
	fmt.Printf("Request method is: %s\n", r.Method)
}

func headerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

// writing a function to set the content-type of the header
func setContentTypePlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Content-type set to plain..")
	w.Header().Set("Content-type", "text/plain")
	fmt.Fprint(w, "<h1>Hello Ahsan!</h1>")
}

func timeoutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Attempting Timeout")
	/*
		We added a timeout for one second in the configurations
		But we slept the method for 2 seconds
		By this way it will not be able to write anything to ResponseWriter
	*/
	time.Sleep(time.Second * 2)
	fmt.Fprint(w, "Did not timeout")
}

func HelloWorldNinja(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Stealth Mode")
	fmt.Fprint(w, "<h1 style=\"background-color:grey;\">Stealh Hello from Ninja</h1>")
}

func main() {
	// Linking main route prompt differant things depending on situation..
	http.HandleFunc("/", handleFunction)

	// Starting the server..
	// http.ListenAndServe("", nil)

	// writing custom server instance..
	server := http.Server{
		Addr:    "",
		Handler: nil,
		/*
			Additional Server Configurations..
			1. Setting the reader timeout to be 1 second.
			2. Setting the writer timeout to be 1 second.
		*/
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}
	// server.ListenAndServe()

	// Creating our own handler or our own multiplexer..
	var ownMUX http.ServeMux
	// Updating the handler to be ownMUX..
	server.Handler = &ownMUX
	/*
		Ensuring that the multiplexer has
		some routing before starting the server
	*/
	ownMUX.HandleFunc("/", handleFunction)
	// Starting the server afterwards..
	server.ListenAndServe()
}
