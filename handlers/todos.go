package handlers

import (
	"errors"
	"log"
	"net/http"
	"text/template"

	"github.com/wilburhimself/todo_go/database"
	"github.com/wilburhimself/todo_go/models"
	"github.com/wilburhimself/todo_go/types"
)

type key int

const todoIDKey key = 0

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	todos := []models.Todo{}
	db.Order("id desc").Find(&todos)

	data := map[string][]models.Todo{
		"todos": todos,
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/todo-item.html"))
	tmpl.Execute(w, data)
}

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	title := r.FormValue("title")

	todo := models.Todo{
		Title: title,
		Done:  false,
	}

	db.Create(&todo)

	tmpl := template.Must(template.ParseFiles("templates/todo-item.html"))
	tmpl.ExecuteTemplate(w, "todo-item", todo)
}

func GetTodoID(r *http.Request) (string, error) {
	todoIDVal := r.Context().Value(types.TodoIDKey)
	if todoIDVal == nil {
		return "", errors.New("todoID not found in context")
	}

	todoID, ok := todoIDVal.(string)
	if !ok {
		return "", errors.New("todoID is not a string")
	}

	return todoID, nil
}

func ToggleTodoHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	todoID, err := GetTodoID(r)

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

func EditTodoHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	todoID, err := GetTodoID(r)
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

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	todoID, err := GetTodoID(r)
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

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	todoID, err := GetTodoID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo := models.Todo{}
	db.First(&todo, todoID)

	db.Delete(&todo)

	tmpl := template.Must(template.ParseFiles("templates/todo-item.html"))
	tmpl.ExecuteTemplate(w, "todo-item", todo)
}

func CancelEditHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	todoID, err := GetTodoID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo := models.Todo{}
	db.First(&todo, todoID)

	tmpl := template.Must(template.ParseFiles("templates/todo-item.html"))
	tmpl.ExecuteTemplate(w, "todo-item", todo)
}
