package main

import (
	"fmt"
	"log"

	"github.com/wilburhimself/todo_go/database"
)

func main() {
	if err := database.InitDB("dev.db"); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	fmt.Println("Database initialized successfully.")
}
