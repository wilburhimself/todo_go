package lib

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/wilburhimself/todo_go/handlers"
	"github.com/wilburhimself/todo_go/types"
)

func Router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", handlers.IndexHandler)
	r.Route("/todos", func(r chi.Router) {
		r.Post("/add", handlers.AddTodoHandler)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(TodoCtx)
			r.Post("/toggle", handlers.ToggleTodoHandler)
			r.Get("/edit", handlers.EditTodoHandler)
			r.Post("/update", handlers.UpdateTodoHandler)
			r.Delete("/delete", handlers.DeleteTodoHandler)
			r.Get("/cancel", handlers.CancelEditHandler)
		})
	})

	return r
}

func TodoCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todoID := chi.URLParam(r, "id")

		ctx := context.WithValue(r.Context(), types.TodoIDKey, todoID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
