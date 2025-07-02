package taskService

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	WhatIsTheTask string `json:"what_is_the_task"`
	IsDone        bool   `json:"is_done"`
}
