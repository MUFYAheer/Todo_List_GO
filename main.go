package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Todo interface
type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Completed   bool   `json:"completed,omitempty"`
}

var todos []Todo

func addTodoEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func listTodosEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func completeTodoEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = true
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

func main() {
	r := mux.NewRouter()
	todos = append(todos, Todo{
		ID: "1", Title: "Go to Gym",
		Description: "You are getting so fat",
		Completed:   false,
	})
	todos = append(todos, Todo{
		ID: "2", Title: "Go for walk",
		Description: "You are getting so fat",
		Completed:   false,
	})
	r.HandleFunc("/", addTodoEndpoint).Methods("POST")
	r.HandleFunc("/", listTodosEndpoint).Methods("GET")
	r.HandleFunc("/{id}", completeTodoEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":1234", r))
}
