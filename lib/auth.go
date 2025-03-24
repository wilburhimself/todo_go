package lib

import (
	"context"
	"net/http"

	"github.com/wilburhimself/todo_go/database"
	"github.com/wilburhimself/todo_go/models"
	"github.com/wilburhimself/todo_go/session"
	"github.com/wilburhimself/todo_go/types"
)

// AuthMiddleware checks if the user is authenticated
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if user is authenticated
		session, exists := session.GetSession(r)
		if !exists || session.UserID == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Get user from database
		db := database.GetDB()
		var user models.User
		if result := db.First(&user, session.UserID); result.Error != nil {
			// User not found, clear session and redirect to login
			// session.ClearSession(w, r)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Add user to context
		ctx := context.WithValue(r.Context(), types.UserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
