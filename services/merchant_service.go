package services

import (
	. "go-project/entities"
	. "go-project/repositories"
)

type MerchantService interface {
	FindAll() ([]Merchant, error)
	FindOne(id string) (*Merchant, error)
	Create(input CreateMerchantInput) (*Merchant, error)
	Update(id string, MerchantInput *UpdateMerchantInput) (*Merchant, error)
	Delete(id string) (*Merchant, error)
}

type merchantService struct {
	repository MerchantRepository
}

func NewMerchantService(r MerchantRepository) *merchantService {
	return &merchantService{r}
}

func (s *merchantService) FindAll() ([]Merchant, error) {
	merchants, err := s.repository.FindAll()
	return merchants, err
}

func (s *merchantService) FindOne(id string) (*Merchant, error) {
	merchant, err := s.repository.FindOne(id)
	return merchant, err
}

func (s *merchantService) Create(input CreateMerchantInput) (*Merchant, error) {
	merchantInput := Merchant{
		Name:       input.Name,
		SocialLink: input.SocialLink,
		UserID:     input.UserID,
	}
	merchant, err := s.repository.Create(&merchantInput)
	return merchant, err
}

func (s *merchantService) Update(id string, merchantInput *UpdateMerchantInput) (*Merchant, error) {
	merchant, err := s.repository.FindOne(id)
	if err != nil {
		return merchant, err
	}
	merchant, err = s.repository.Update(merchant, merchantInput)
	return merchant, err
}

func (s *merchantService) Delete(id string) (*Merchant, error) {
	merchant, err := s.repository.FindOne(id)
	if err != nil {
		return merchant, err
	}
	merchant, err = s.repository.Delete(merchant)
	return merchant, err
}
