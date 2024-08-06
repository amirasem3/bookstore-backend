package usecase

import (
	"bookstore/internal/entity"
	"bookstore/internal/repository"
	"errors"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type TaskUsecase interface {
	GetTasks() ([]*entity.Task, error)
	GetTaskByID(id int) (*entity.Task, error)
	CreateTask(task *entity.Task) error
	UpdateTask(task *entity.Task) error
	DeleteTask(id int) error
}

type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{repo: repo}
}

func (u *taskUsecase) GetTasks() ([]*entity.Task, error) {
	return u.repo.GetTasks()
}

func (u *taskUsecase) GetTaskByID(id int) (*entity.Task, error) {
	return u.repo.GetTaskById(id)
}

func (u *taskUsecase) CreateTask(task *entity.Task) error {
	return u.repo.CreateTask(task)
}

func (u *taskUsecase) UpdateTask(task *entity.Task) error {
	return u.repo.UpdateTask(task)
}
func (u *taskUsecase) DeleteTask(id int) error {
	return u.repo.DeleteTask(id)
}
