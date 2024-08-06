package repository

import (
	"bookstore/internal/entity"
	"database/sql"
	"errors"
	"log"
)

type TaskRepository interface {
	GetTasks() ([]*entity.Task, error)
	GetTaskById(id int) (*entity.Task, error)
	CreateTask(task *entity.Task) error
	UpdateTask(task *entity.Task) error
	DeleteTask(id int) error
}

type sqlTaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &sqlTaskRepository{db: db}
}

func (r *sqlTaskRepository) GetTasks() ([]*entity.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description, completed FROM tasks")
	if err != nil {
		log.Println("Error querying tasks:", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []*entity.Task
	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
		if err != nil {
			log.Println("Error scanning task:", err)
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over task rows:", err)
		return nil, err
	}

	return tasks, nil
}

func (r *sqlTaskRepository) GetTaskById(id int) (*entity.Task, error) {
	var task entity.Task
	err := r.db.QueryRow("SELECT id, title, description, completed FROM tasks WHERE id = @p1", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("task not found")
		}
		log.Println("Error querying task by ID:", err)
		return nil, err
	}
	return &task, nil
}

func (r *sqlTaskRepository) CreateTask(task *entity.Task) error {

	query := `
		INSERT INTO tasks (title, description, completed) 
		VALUES (@p1, @p2, @p3);
		SELECT ID = SCOPE_IDENTITY();
	`
	// Use QueryRow to fetch the inserted ID using SCOPE_IDENTITY()
	err := r.db.QueryRow(query, task.Title, task.Description, task.Completed).Scan(&task.ID)
	if err != nil {
		log.Println("Error inserting task:", err)
		return err
	}

	return nil
}

func (r *sqlTaskRepository) UpdateTask(task *entity.Task) error {
	_, err := r.db.Exec("UPDATE tasks SET title = @p1, description = @p2, completed = @p3 WHERE id = @p4",
		task.Title, task.Description, task.Completed, task.ID)
	if err != nil {
		log.Println("Error updating task:", err)
		return err
	}
	return nil
}

func (r *sqlTaskRepository) DeleteTask(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = @p1", id)
	if err != nil {
		log.Println("Error deleting task:", err)
		return err
	}
	return nil
}

//type inMemoryTaskRepository struct {
//	tasks  map[int]*entity.Task
//	nextID int
//}

//func NewTaskRepository() TaskRepository {
//	return &inMemoryTaskRepository{
//		tasks:  make(map[int]*entity.Task),
//		nextID: 1,
//	}
//}

//func (repo *inMemoryTaskRepository) GetTasks() ([]*entity.Task, error) {
//	var tasklist []*entity.Task
//	for _, task := range repo.tasks {
//		tasklist = append(tasklist, task)
//	}
//	return tasklist, nil
//}
//
//func (repo *inMemoryTaskRepository) GetTaskById(id int) (*entity.Task, error) {
//	if task, exists := repo.tasks[id]; exists {
//		return task, nil
//	}
//	return nil, errors.New("task not found")
//}
//
//func (repo *inMemoryTaskRepository) CreateTask(task *entity.Task) error {
//	task.ID = repo.nextID
//	repo.tasks[repo.nextID] = task
//	repo.nextID++
//	return nil
//}
//
//func (repo *inMemoryTaskRepository) UpdateTask(task *entity.Task) error {
//	if _, exists := repo.tasks[task.ID]; exists {
//		repo.tasks[task.ID] = task
//		return nil
//	}
//	return errors.New("task not found")
//}
//
//func (repo *inMemoryTaskRepository) DeleteTask(id int) error {
//	if _, exists := repo.tasks[id]; exists {
//		delete(repo.tasks, id)
//		return nil
//	}
//	return errors.New("task not found")
//}
