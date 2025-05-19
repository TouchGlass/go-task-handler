package service

import (
	"errors"
	"strings"
)

type TaskService interface {
	CreateTask(task Task) (error, Task)
	GetTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id string, task Task) (Task, error)
	DeleteTaskByID(id string) error
}
type taskService struct {
	repo TaskRepository
}

// создание сервиса
func NewTaskService(repo TaskRepository) *taskService {
	return &taskService{repo: repo}
}

func (ts taskService) TaskCheck(task Task) error {
	trimmed := strings.TrimSpace(task.WhatIsTheTask)

	if trimmed == "" {
		return errors.New("task title cannot be empty")
	}

	if len(trimmed) > 200 {

		return errors.New("task title is too long (maximum is 200 characters)")
	}

	return nil
}

func (ts taskService) CreateTask(task Task) (error, Task) {
	if err := ts.TaskCheck(task); err != nil {
		return err, Task{}
	}
	err := ts.repo.CreateTask(task)
	if err != nil {
		return err, Task{}
	}
	return nil, task
}

func (ts taskService) GetTasks() ([]Task, error) {
	return ts.repo.GetTasks()
}

func (ts taskService) GetTaskByID(id string) (Task, error) {
	return ts.repo.GetTaskByID(id)
}

func (ts taskService) UpdateTask(id string, task Task) (Task, error) {

	var dbtask Task
	dbtask, err := ts.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	dbtask.WhatIsTheTask = task.WhatIsTheTask
	dbtask.IsDone = task.IsDone

	if err := ts.repo.UpdateTask(dbtask); err != nil {
		return Task{}, err
	}

	return dbtask, nil
}

func (ts taskService) DeleteTaskByID(id string) error {
	return ts.repo.DeleteTaskByID(id)
}
