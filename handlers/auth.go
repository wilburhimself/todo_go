package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/wilburhimself/todo_go/database"
	"github.com/wilburhimself/todo_go/models"
	"github.com/wilburhimself/todo_go/session"
	"gorm.io/gorm"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	// Check if user is already logged in
	session, exists := session.GetSession(r)
	if exists && session.UserID != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"Error": r.URL.Query().Get("error"),
	}

	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, data)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Redirect(w, r, "/login?error=Please+provide+username+and+password", http.StatusSeeOther)
		return
	}

	db := database.GetDB()
	var user models.User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Redirect(w, r, "/login?error=Invalid+username+or+password", http.StatusSeeOther)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Check password
	if !user.CheckPassword(password) {
		http.Redirect(w, r, "/login?error=Invalid+username+or+password", http.StatusSeeOther)
		return
	}

	// Create session
	session.CreateSession(w, user.ID)

	// Redirect to homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	// Check if user is already logged in
	session, exists := session.GetSession(r)
	if exists && session.UserID != 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"Error": r.URL.Query().Get("error"),
	}

	tmpl := template.Must(template.ParseFiles("templates/register.html"))
	tmpl.Execute(w, data)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	if username == "" || email == "" || password == "" {
		http.Redirect(w, r, "/register?error=All+fields+are+required", http.StatusSeeOther)
		return
	}

	if password != confirmPassword {
		http.Redirect(w, r, "/register?error=Passwords+do+not+match", http.StatusSeeOther)
		return
	}

	db := database.GetDB()

	// Check if username already exists
	var count int64
	db.Model(&models.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		http.Redirect(w, r, "/register?error=Username+already+exists", http.StatusSeeOther)
		return
	}

	// Check if email already exists
	db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		http.Redirect(w, r, "/register?error=Email+already+exists", http.StatusSeeOther)
		return
	}

	// Create user
	user := models.User{
		Username: username,
		Email:    email,
	}

	// Set password - this will be hashed by the model's SetPassword method
	if err := user.SetPassword(password); err != nil {
		http.Error(w, "Failed to process password", http.StatusInternalServerError)
		return
	}

	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Create session
	session.CreateSession(w, user.ID)

	// Redirect to homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session.ClearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
