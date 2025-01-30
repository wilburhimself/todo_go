package main

import (
	"log"
	"net/http"

	"github.com/wilburhimself/todo_go/database"
	"github.com/wilburhimself/todo_go/lib"
)

func main() {
	if err := database.InitDB("dev.db"); err != nil {
		log.Fatal(err)
	}

	router := lib.Router()
	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
