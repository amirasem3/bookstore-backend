package httpHandler

import (
	"bookstore/internal/entity"
	"bookstore/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	usecase usecase.TaskUsecase
}

func NewTaskHandler(usecase usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{usecase: usecase}
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getTasks(w, r)
	case http.MethodPost:
		h.createTask(w, r)
	case http.MethodPut:
		h.updateTask(w, r)
	case http.MethodDelete:
		h.deleteTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// New handler for updating the completed status
func (h *TaskHandler) UpdateCompletedStatus(w http.ResponseWriter, r *http.Request) {
	var data struct {
		ID        int  `json:"id"`
		Completed bool `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.usecase.UpdateTaskStatus(data.ID, data.Completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}

func (h *TaskHandler) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.usecase.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var task entity.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.usecase.CreateTask(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task) // Return the created task with all fields

}

func (h *TaskHandler) updateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.usecase.UpdateTask(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task) // Return the created task with all fields

}

func (h *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}
	if err := h.usecase.DeleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}
