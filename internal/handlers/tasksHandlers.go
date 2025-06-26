package handlers

import (
	"BDproj/internal/service"
	"BDproj/internal/web/tasks"
	"context"
	"fmt"
)

type Handler struct {
	service service.TaskService
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:            &tsk.ID,
			IsDone:        &tsk.IsDone,
			WhatIsTheTask: tsk.WhatIsTheTask,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {

	taskRequest := request.Body

	taskToCreate := service.Task{
		WhatIsTheTask: taskRequest.WhatIsTheTask,
		IsDone:        *taskRequest.IsDone,
	}

	err, createdTask := h.service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:            &createdTask.ID,
		WhatIsTheTask: createdTask.WhatIsTheTask,
		IsDone:        &createdTask.IsDone,
	}
	return response, nil
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskIDstr := fmt.Sprintf("%d", request.Id)

	if err := h.service.DeleteTaskByID(taskIDstr); err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil

}

func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskIDstr := fmt.Sprintf("%d", request.Id)

	patchBody := request.Body

	taskToUpdate := service.Task{
		WhatIsTheTask: *patchBody.WhatIsTheTask,
		IsDone:        *patchBody.IsDone,
	}

	task, err := h.service.UpdateTask(taskIDstr, taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:            &task.ID,
		WhatIsTheTask: task.WhatIsTheTask,
		IsDone:        &task.IsDone,
	}

	return response, err
}

// создание хэндлеров
func NewHandler(s service.TaskService) *Handler {
	return &Handler{service: s}
}
