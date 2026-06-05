package main

import (
	"fmt"
	"net/http"
	"os"

	"go-supabase-todos/db"
	"go-supabase-todos/handler"
	"go-supabase-todos/repository"
	"go-supabase-todos/service"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	godotenv.Load()

	// connect to supabase
	database := db.Connect()
	defer database.Close()

	// wire up layers
	repo := repository.NewPostgresRepo(database)
	svc := service.NewTodoService(repo)
	h := handler.NewTodoHandler(svc)

	// routes
	http.HandleFunc("GET /todos", h.GetTodos)
	http.HandleFunc("POST /todos", h.CreateTodo)
	http.HandleFunc("GET /todos/{id}", h.GetTodo)
	http.HandleFunc("PUT /todos/{id}", h.UpdateTodo)
	http.HandleFunc("DELETE /todos/{id}", h.DeleteTodo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4321" // fallback for local
	}
	
	fmt.Printf("server running on :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
