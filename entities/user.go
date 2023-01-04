package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string
	Email    string
}

type CreateUserInput struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UpdateUserInput struct {
	gorm.Model
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
