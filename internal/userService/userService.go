package userService

import (
	"BDproj/internal/taskService"
	"errors"
)

type UserService interface {
	CreateUser(user User) (error, User)
	GetUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(id string, user User) (User, error)
	DeleteUserByID(id string) error
	GetTasksForUser(userID uint) ([]taskService.Task, error)
}
type userService struct {
	userRepo UserRepository
}

// создание сервиса
func NewUserService(userRepo UserRepository) *userService {
	return &userService{userRepo: userRepo}
}

func (us userService) PasswordCheck(user User) error {

	if len(user.Password) < 4 {
		return errors.New("password must be at least 4 characters long")
	}

	return nil
}

func (us userService) CreateUser(user User) (error, User) {
	if err := us.PasswordCheck(user); err != nil {
		return err, User{}
	}
	createduser, err := us.userRepo.CreateUser(user)
	if err != nil {
		return err, User{}
	}
	return nil, createduser
}

func (us userService) GetUsers() ([]User, error) {
	return us.userRepo.GetUsers()
}

func (us userService) GetUserByID(id string) (User, error) {
	return us.userRepo.GetUserByID(id)
}

func (us userService) UpdateUser(id string, user User) (User, error) {

	var dbuser User
	dbuser, err := us.GetUserByID(id)
	if err != nil {
		return User{}, err
	}

	dbuser.Email = user.Email
	dbuser.Password = user.Password

	if err := us.userRepo.UpdateUser(dbuser); err != nil {
		return User{}, err
	}

	return dbuser, nil
}

func (us userService) DeleteUserByID(id string) error {
	return us.userRepo.DeleteUserByID(id)
}

func (us userService) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	tasks, err := us.userRepo.GetTasksForUser(userID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
