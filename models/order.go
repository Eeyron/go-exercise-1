package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Status     string
	UserID     uint
	User       User
	OrderItems []OrderItem
}

type CreateOrderInput struct {
	Status                string                 `json:"status" binding:"required"`
	UserID                uint                   `json:"user_id" binding:"required"`
	CreateOrderItemInputs []CreateOrderItemInput `json:"order_items" binding:"required"`
}

type UpdateOrderInput struct {
	gorm.Model
	Status string `json:"status"`
	UserID uint   `json:"user_id"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	Quantity  uint
	ProductID uint
	Product   Product
}

type CreateOrderItemInput struct {
	Quantity  uint `json:"quantity" binding:"required"`
	ProductID uint `json:"product_id" binding:"required"`
}

type UpdateOrderItemInput struct {
	gorm.Model
	OrderID   uint `json:"order_id"`
	Quantity  uint `json:"quantity"`
	ProductID uint `json:"product_id"`
}
