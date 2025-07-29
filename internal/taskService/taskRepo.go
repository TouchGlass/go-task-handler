package taskService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(task Task) error
	DeleteTaskByID(id string) error
}

type taskRepository struct {
	db *gorm.DB
}

// создание репы
func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (repo *taskRepository) CreateTask(task Task) (Task, error) {
	err := repo.db.Create(&task).Error
	return task, err
}

func (repo *taskRepository) GetTasks() ([]Task, error) {
	var tasks []Task
	err := repo.db.Find(&tasks).Error
	return tasks, err
}

func (repo *taskRepository) GetTaskByID(id string) (Task, error) {
	var task Task
	err := repo.db.First(&task, "id = ?", id).Error
	return task, err
}

func (repo *taskRepository) UpdateTask(task Task) error {
	return repo.db.Save(&task).Error
}

func (repo *taskRepository) DeleteTaskByID(id string) error {
	return repo.db.Delete(&Task{}, "id = ?", id).Error
}
