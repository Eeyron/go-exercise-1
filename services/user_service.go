package services

import (
	. "go-project/entities"
	. "go-project/repositories"
)

type UserService interface {
	FindAll() ([]User, error)
	FindOne(id string) (*User, error)
	Create(input CreateUserInput) (*User, error)
	Update(id string, userInput *UpdateUserInput) (*User, error)
	Delete(id string) (*User, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(r UserRepository) *userService {
	return &userService{r}
}

func (s *userService) FindAll() ([]User, error) {
	users, err := s.repository.FindAll()
	return users, err
}

func (s *userService) FindOne(id string) (*User, error) {
	user, err := s.repository.FindOne(id)
	return user, err
}

func (s *userService) Create(input CreateUserInput) (*User, error) {
	userInput := User{
		FullName: input.FullName,
		Email:    input.Email,
	}
	user, err := s.repository.Create(&userInput)
	return user, err
}

func (s *userService) Update(id string, userInput *UpdateUserInput) (*User, error) {
	user, err := s.repository.FindOne(id)
	if err != nil {
		return user, err
	}
	user, err = s.repository.Update(user, userInput)
	return user, err
}

func (s *userService) Delete(id string) (*User, error) {
	user, err := s.repository.FindOne(id)
	if err != nil {
		return user, err
	}
	user, err = s.repository.Delete(user)
	return user, err
}
