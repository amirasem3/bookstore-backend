package repository

import (
	"bookstore/internal/entity"
	"errors"
)

type TaskRepository interface {
	GetTasks() ([]*entity.Task, error)
	GetTaskById(id int) (*entity.Task, error)
	CreateTask(task *entity.Task) error
	UpdateTask(task *entity.Task) error
	DeleteTask(id int) error
}

type inMemoryTaskRepository struct {
	tasks  map[int]*entity.Task
	nextID int
}

func NewTaskRepository() TaskRepository {
	return &inMemoryTaskRepository{
		tasks:  make(map[int]*entity.Task),
		nextID: 1,
	}
}

func (repo *inMemoryTaskRepository) GetTasks() ([]*entity.Task, error) {
	var tasklist []*entity.Task
	for _, task := range repo.tasks {
		tasklist = append(tasklist, task)
	}
	return tasklist, nil
}

func (repo *inMemoryTaskRepository) GetTaskById(id int) (*entity.Task, error) {
	if task, exists := repo.tasks[id]; exists {
		return task, nil
	}
	return nil, errors.New("task not found")
}

func (repo *inMemoryTaskRepository) CreateTask(task *entity.Task) error {
	task.ID = repo.nextID
	repo.tasks[repo.nextID] = task
	repo.nextID++
	return nil
}

func (repo *inMemoryTaskRepository) UpdateTask(task *entity.Task) error {
	if _, exists := repo.tasks[task.ID]; exists {
		repo.tasks[task.ID] = task
		return nil
	}
	return errors.New("task not found")
}

func (repo *inMemoryTaskRepository) DeleteTask(id int) error {
	if _, exists := repo.tasks[id]; exists {
		delete(repo.tasks, id)
		return nil
	}
	return errors.New("task not found")
}
