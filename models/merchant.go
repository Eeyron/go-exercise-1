package models

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	Name       string
	SocialLink string
	UserID     uint
	User       User
}

type CreateMerchantInput struct {
	Name       string `json:"name" binding:"required"`
	SocialLink string `json:"social_link"`
	UserID     uint   `json:"user_id" binding:"required"`
}

type UpdateMerchantInput struct {
	gorm.Model
	Name       string `json:"name"`
	SocialLink string `json:"social_link"`
	UserID     uint   `json:"user_id"`
}
