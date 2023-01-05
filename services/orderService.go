package services

import (
	. "go-project/entities"
	. "go-project/repositories"
)

type OrderService interface {
	FindAll() ([]Order, error)
	FindOne(id string) (*Order, error)
	Create(input CreateOrderInput) (*Order, error)
}

type orderService struct {
	repository OrderRepository
}

func NewOrderService(r OrderRepository) *orderService {
	return &orderService{r}
}

func (s *orderService) FindAll() ([]Order, error) {
	orders, err := s.repository.FindAll()
	return orders, err
}

func (s *orderService) FindOne(id string) (*Order, error) {
	order, err := s.repository.FindOne(id)
	return order, err
}

func (s *orderService) Create(input CreateOrderInput) (*Order, error) {
	orderInput := Order{
		Status: input.Status,
		UserID: input.UserID,
	}
	order, err := s.repository.Create(&orderInput, input.CreateOrderItemInputs)
	return order, err
}
