package main

import (
	"errors"
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

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
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

func getTodoID(r *http.Request) (uint, error) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		return 0, errors.New("invalid path")
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

func toggleTodoHandler(w http.ResponseWriter, r *http.Request) {
	db := returnDB()

	todoID, err := getTodoID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

func editTodoHandler(w http.ResponseWriter, r *http.Request) {
	db := returnDB()

	todoID, err := getTodoID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo := models.Todo{}
	db.First(&todo, todoID)

	log.Println(todo)

	tmpl := template.Must(template.ParseFiles("templates/edit-item.html"))
	tmpl.ExecuteTemplate(w, "edit-item", todo)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	db := returnDB()

	todoID, err := getTodoID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo := models.Todo{}
	db.First(&todo, todoID)

	log.Println(todo)

	title := r.FormValue("title")
	todo.Title = title

	db.Save(&todo)

	tmpl := template.Must(template.ParseFiles("templates/todo-item.html"))
	tmpl.ExecuteTemplate(w, "todo-item", todo)
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
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
		toggleTodoHandler(w, r)
	case action == "edit" && r.Method == "GET":
		editTodoHandler(w, r)
	case action == "update" && r.Method == "POST":
		updateTodoHandler(w, r)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func main() {
	initDB()

	http.HandleFunc("/", dashboardHandler)
	http.HandleFunc("/add", addTodoHandler)
	http.HandleFunc("/todos/", todosHandler)
	http.ListenAndServe(":8080", nil)
}
