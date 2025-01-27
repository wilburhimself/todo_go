package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/wilburhimself/go_todo_list/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Todo{})
}

func returnDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	db := returnDB()

	todos := []models.Todo{}
	db.Order("id desc").Find(&todos)

	data := map[string][]models.Todo{
		"todos": todos,
	}

	tmpl := template.Must(template.ParseFiles("index.html", "templates/todo-item.html"))
	tmpl.Execute(w, data)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	db := returnDB()

	title := r.FormValue("title")

	todo := models.Todo{
		Title: title,
		Done:  false,
	}

	db.Create(&todo)

	tmpl := template.Must(template.ParseFiles("templates/todo-item.html"))
	tmpl.ExecuteTemplate(w, "todo-item", todo)
}

func toggleTodoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("toggleTodoHandler")
	db := returnDB()

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 || parts[3] != "toggle" {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	todoID, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo := models.Todo{}
	db.First(&todo, todoID)

	log.Println(todo)

	todo.Done = !todo.Done
	db.Save(&todo)

	tmpl := template.Must(template.ParseFiles("templates/todo-item.html"))
	tmpl.ExecuteTemplate(w, "todo-item", todo)
}

func main() {
	initDB()

	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/add", addTodoHandler)
	http.HandleFunc("/todos/", toggleTodoHandler)
	http.ListenAndServe(":8080", nil)
}
