package service

import (
	"errors"
	"strings"
)

type TaskService interface {
	CreateTask(task Task) error
	GetTasks() ([]Task, error)
	PatchTask(id string, task Task) error
	DeleteTaskByID(id string) error
}
type taskService struct {
	repo TaskRepository
}

// создание сервиса
func NewTaskService(repo TaskRepository) *taskService {
	return &taskService{repo: repo}
}

func (ts taskService) CreateTask(task Task) error {
	trimmed := strings.TrimSpace(task.WhatIsTheTask)

	if trimmed == "" {
		return errors.New("task title cannot be empty")
	}

	if len(trimmed) > 200 {
		return errors.New("task title is too long (maximum is 200 characters)")
	}

	// если всё ок — передаём задачу в репозиторий
	return ts.repo.CreateTask(task)
}

func (ts taskService) GetTasks() ([]Task, error) {
	return ts.repo.GetTasks()
}

func (ts taskService) PatchTask(id string, task Task) error {

	var dbtask Task
	dbtask, err := ts.repo.GetTaskByID(id)
	if err != nil {
		return err
	}

	dbtask.WhatIsTheTask = task.WhatIsTheTask
	dbtask.IsDone = task.IsDone

	if err := ts.repo.UpdateTask(dbtask); err != nil {
		return err
	}

	return nil
}

func (ts taskService) DeleteTaskByID(id string) error {
	return ts.repo.DeleteTaskByID(id)
}
