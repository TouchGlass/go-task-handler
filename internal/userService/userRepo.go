package userService

import (
	"BDproj/internal/taskService"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(user User) error
	DeleteUserByID(id string) error
	GetTasksForUser(userID uint) ([]taskService.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

// создание репы
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(user User) (User, error) {
	return user, repo.db.Create(&user).Error
}

func (repo *userRepository) GetUsers() ([]User, error) {
	var users []User
	err := repo.db.Find(&users).Error
	return users, err
}

func (repo *userRepository) GetUserByID(id string) (User, error) {
	var user User
	err := repo.db.First(&user, "id = ?", id).Error
	return user, err
}

func (repo *userRepository) UpdateUser(user User) error {
	return repo.db.Save(&user).Error
}

func (repo *userRepository) DeleteUserByID(id string) error {
	return repo.db.Unscoped().Delete(&User{}, "id = ?", id).Error
}

func (repo *userRepository) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	var tasks []taskService.Task
	err := repo.db.
		Where("user_id = ?", userID).
		Find(&tasks).Error
	return tasks, err
}
