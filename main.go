package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string = "World"

type TaskRequest struct {
	Task string 'json:"task"'
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var taskreq TaskRequest
	err := json.NewDecoder(r.Body).Decode(&taskreq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	task = taskreq.Task
	fmt.Fprintln(w, "Task updated successfully!")
}


func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", task)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/task", TaskHandler).Methods("POST")
	router.HandleFunc("/hello", HelloHandler).Methods("GET")
	http.ListenAndServe(":8080", router)

}
