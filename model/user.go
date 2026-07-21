package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	Age  int    `json:"age"`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required,min=1,max=120"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age" validate:"min=1,max=120"`
}