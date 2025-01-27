package main

import (
	"net/http"
	"strings"

	"github.com/wilburhimself/todo_go/handlers"
)

func TodosHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	action := ""
	if len(parts) >= 4 {
		action = parts[3]
	}

	switch {
	case action == "toggle" && r.Method == "POST":
		handlers.ToggleTodoHandler(w, r)
	case action == "edit" && r.Method == "GET":
		handlers.EditTodoHandler(w, r)
	case action == "update" && r.Method == "POST":
		handlers.UpdateTodoHandler(w, r)
	case action == "delete" && r.Method == "DELETE":
		handlers.DeleteTodoHandler(w, r)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/add", handlers.AddTodoHandler)
	http.HandleFunc("/todos/", TodosHandler)
	http.ListenAndServe(":8080", nil)
}
