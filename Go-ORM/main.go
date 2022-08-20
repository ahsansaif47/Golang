package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", rootPage).Methods("GET")
	router.HandleFunc("/users", allUsers).Methods("GET")
	router.HandleFunc("/users/{name}/{email}", newUser).Methods("POST")
	router.HandleFunc("/users/{name}", deleteUser).Methods("DELETE")
	router.HandleFunc("/users/{name}/{email}", updateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	fmt.Println("Go ORM Tutorial")
	// Initial migration before using the API
	initialMigration()
	handleRequests()
}
