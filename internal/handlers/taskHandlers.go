package handlers

import (
	"BDproj/internal/taskService"
	"BDproj/internal/web/tasks"
	"context"
	"fmt"
)

type TaskHandler struct {
	taskService taskService.TaskService
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.taskService.GetTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:            &tsk.ID,
			IsDone:        &tsk.IsDone,
			WhatIsTheTask: tsk.WhatIsTheTask,
			UserId:        tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		WhatIsTheTask: taskRequest.WhatIsTheTask,
		IsDone:        *taskRequest.IsDone,
		UserID:        taskRequest.UserId,
	}

	err, createdTask := h.taskService.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:            &createdTask.ID,
		WhatIsTheTask: createdTask.WhatIsTheTask,
		IsDone:        &createdTask.IsDone,
		UserId:        createdTask.UserID,
	}
	return response, nil
}

func (h *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskIDstr := fmt.Sprintf("%d", request.Id)

	if err := h.taskService.DeleteTaskByID(taskIDstr); err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil

}

func (h *TaskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskIDstr := fmt.Sprintf("%d", request.Id)

	patchBody := request.Body

	taskToUpdate := taskService.Task{
		WhatIsTheTask: *patchBody.WhatIsTheTask,
		IsDone:        *patchBody.IsDone,
		UserID:        *patchBody.UserId,
	}

	task, err := h.taskService.UpdateTask(taskIDstr, taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:            &task.ID,
		WhatIsTheTask: task.WhatIsTheTask,
		IsDone:        &task.IsDone,
		UserId:        *patchBody.UserId,
	}

	return response, err
}

// создание хэндлеров
func NewTaskHandler(ts taskService.TaskService) *TaskHandler {
	return &TaskHandler{taskService: ts}
}
