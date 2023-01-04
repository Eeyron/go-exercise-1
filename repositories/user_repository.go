package repositories

import (
	. "go-project/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]User, error)
	FindOne(id string) (*User, error)
	Create(userInput *User) (*User, error)
	Update(user *User, userInput *UpdateUserInput) (*User, error)
	Delete(user *User) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepository) FindOne(id string) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *User) (*User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Update(user *User, userInput *UpdateUserInput) (*User, error) {
	if err := r.db.Model(&user).Updates(userInput).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Delete(user *User) (*User, error) {
	if err := r.db.Delete(user).Error; err != nil {
		return user, err
	}
	return user, nil
}
