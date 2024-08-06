package main

import (
	"bookstore/internal/config"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"path/filepath"

	"bookstore/internal/delivery/httpHandler"
	"bookstore/internal/repository"
	"bookstore/internal/usecase"
)

func main() {

	absPath := filepath.Join("C:\\", "Users", "AH.Yousefi", "GolandProjects", "bookstore-backend", ".env")
	err := godotenv.Load(absPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database connection
	db, err := config.GetDBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	//Initialize Repository
	repo := repository.NewTaskRepository(db)

	//Initialize use case
	taskUseCase := usecase.NewTaskUsecase(repo)

	//Initialize task handler
	taskHandler := httpHandler.NewTaskHandler(taskUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", taskHandler.HandleTasks)
	mux.HandleFunc("/tasks/update-status", taskHandler.UpdateCompletedStatus) // New route for updating status

	corsHandler := httpHandler.CORS(mux)

	//http.HandleFunc("/tasks", taskHandler.HandleTasks)
	log.Println("server starter at port 8080")
	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
