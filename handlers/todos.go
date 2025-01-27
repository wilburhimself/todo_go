package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/wilburhimself/todo_go/lib"
	"github.com/wilburhimself/todo_go/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	db := lib.ReturnDB()

	todos := []models.Todo{}
	db.Order("id desc").Find(&todos)

	data := map[string][]models.Todo{
		"todos": todos,
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/todo-item.html"))
	tmpl.Execute(w, data)
}

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	db := lib.ReturnDB()

	title := r.FormValue("title")

	todo := models.Todo{
		Title: title,
		Done:  false,
	}

	db.Create(&todo)

	tmpl := template.Must(template.ParseFiles("templates/todo-item.html"))
	tmpl.ExecuteTemplate(w, "todo-item", todo)
}

func GetTodoID(r *http.Request) (uint, error) {
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

func ToggleTodoHandler(w http.ResponseWriter, r *http.Request) {
	db := lib.ReturnDB()

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
	db := lib.ReturnDB()

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
	db := lib.ReturnDB()

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
	db := lib.ReturnDB()

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
