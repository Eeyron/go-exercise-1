package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	Price      float64
	Quantity   uint
	MerchantID uint
}

type CreateProductInput struct {
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	Quantity   uint    `json:"quantity" binding:"required"`
	MerchantID uint    `json:"merchant_id" binding:"required"`
}

type UpdateProductInput struct {
	gorm.Model
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Quantity   uint    `json:"quantity"`
	MerchantID uint    `json:"merchant_id"`
}
