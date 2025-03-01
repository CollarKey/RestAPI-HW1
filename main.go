package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Hello World")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler)
	http.ListenAndServe(":8080", router)

	fmt.Println("Server started")
}
