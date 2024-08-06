package main

import (
	"log"
	"net/http"

	"bookstore/internal/delivery/httpHandler"
	"bookstore/internal/repository"
	"bookstore/internal/usecase"
)

func main() {
	//Initialize Repository
	repo := repository.NewTaskRepository()

	//Initialize use case
	taskUseCase := usecase.NewTaskUsecase(repo)

	//Initialize task handler
	taskHandler := httpHandler.NewTaskHandler(taskUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", taskHandler.HandleTasks)

	corsHandler := httpHandler.CORS(mux)

	//http.HandleFunc("/tasks", taskHandler.HandleTasks)
	log.Println("server starter at port 8080")
	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
