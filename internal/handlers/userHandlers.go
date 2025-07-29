package handlers

import (
	userService2 "BDproj/internal/userService"
	"BDproj/internal/web/users"
	"context"
	"fmt"
)

type UserHandler struct {
	userService userService2.UserService
}

func (u UserHandler) GetTasksByUserID(_ context.Context, request users.GetTasksByUserIDRequestObject) (users.GetTasksByUserIDResponseObject, error) {
	userID := request.Id
	userTasks, err := u.userService.GetTasksForUser(userID)
	if err != nil {
		return nil, err
	}
	response := users.GetTasksByUserID200JSONResponse{}

	for _, usertasks := range userTasks {
		task := users.Task{
			Id:            &usertasks.ID,
			WhatIsTheTask: usertasks.WhatIsTheTask,
			IsDone:        &usertasks.IsDone,
			UserId:        usertasks.UserID,
		}
		response = append(response, task)

	}
	return response, err
}

func (u UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.userService.GetUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    usr.Email,
			Password: usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (u UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService2.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	err, createdUser := u.userService.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    createdUser.Email,
		Password: createdUser.Password,
	}
	return response, nil
}

func (u UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userIDstr := fmt.Sprintf("%d", request.Id)

	if err := u.userService.DeleteUserByID(userIDstr); err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil

}

func (u UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userIDstr := fmt.Sprintf("%d", request.Id)

	patchBody := request.Body

	userToUpdate := userService2.User{
		Email:    *patchBody.Email,
		Password: *patchBody.Password,
	}

	user, err := u.userService.UpdateUser(userIDstr, userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &user.ID,
		Email:    user.Email,
		Password: user.Password,
	}

	return response, err
}

// создание хэндлеров
func NewUserHandler(us userService2.UserService) *UserHandler {
	return &UserHandler{userService: us}
}
