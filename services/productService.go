package services

import (
	. "go-project/entities"
	. "go-project/repositories"
)

type ProductService interface {
	FindAll() ([]Product, error)
	FindOne(id string) (*Product, error)
	Create(input CreateProductInput) (*Product, error)
	Update(id string, ProductInput *UpdateProductInput) (*Product, error)
	Delete(id string) (*Product, error)
}

type productService struct {
	repository ProductRepository
}

func NewProductService(r ProductRepository) *productService {
	return &productService{r}
}

func (s *productService) FindAll() ([]Product, error) {
	products, err := s.repository.FindAll()
	return products, err
}

func (s *productService) FindOne(id string) (*Product, error) {
	product, err := s.repository.FindOne(id)
	return product, err
}

func (s *productService) Create(input CreateProductInput) (*Product, error) {
	productInput := Product{
		Name:       input.Name,
		Price:      input.Price,
		Quantity:   input.Quantity,
		MerchantID: input.MerchantID,
	}
	product, err := s.repository.Create(&productInput)
	return product, err
}

func (s *productService) Update(id string, productInput *UpdateProductInput) (*Product, error) {
	product, err := s.repository.FindOne(id)
	if err != nil {
		return product, err
	}
	product, err = s.repository.Update(product, productInput)
	return product, err
}

func (s *productService) Delete(id string) (*Product, error) {
	product, err := s.repository.FindOne(id)
	if err != nil {
		return product, err
	}
	product, err = s.repository.Delete(product)
	return product, err
}
