package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type Articles []Article

func getArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Name: "ABC", Age: "17"},
		Article{Name: "XYZ", Age: "17"},
	}
	json.NewEncoder(w).Encode(articles)
	fmt.Println(articles)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hit on root")
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test Post Articles Endpoint")
}

func handleRequests() {
	// implementing API with net/http library..
	// http.HandleFunc("/", HomePage)
	// http.HandleFunc("/articles", getArticles)
	// http.ListenAndServe("", nil)

	// implementing with gorilla router instance..
	router := mux.NewRouter().StrictSlash(true)
	// advantage of using the gorilla mux library is that we can specify the methods..
	router.HandleFunc("/", HomePage).Methods("GET")
	router.HandleFunc("/articles", getArticles).Methods("GET")
	router.HandleFunc("/articles", testPostArticles).Methods("POST")
	http.ListenAndServe("localhost:8081", router)

}

func main() {
	// http.HandleFunc("/", HomePage)
	// http.HandleFunc("/articles", getArticles)
	// http.ListenAndServe("", nil)
	handleRequests()
}
