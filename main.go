package main

import (
	"encoding/json"
	"fmt"
	"log"
	//"github.com/labstack/echo/v4"
	"net/http"
)

type RequestBody struct {
	Task string `json:"task"`
}

var task []string

func switchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postHandler(w, r)
	case http.MethodGet:
		getHandler(w, r)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Message string   `json:"message"`
		Tasks   []string `json:"tasks"`
	}{
		Message: "hello, ",
		Tasks:   task,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task = append(task, requestBody.Task)
	fmt.Fprintf(w, "Task added successfully: %v", requestBody.Task)
}

func main() {
	http.HandleFunc("/task", switchHandler)

	log.Print("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
