package repositories

import (
	. "go-project/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MerchantRepository interface {
	FindAll() ([]Merchant, error)
	FindOne(id string) (*Merchant, error)
	Create(merchant *Merchant) (*Merchant, error)
	Update(merchant *Merchant, merchantInput *UpdateMerchantInput) (*Merchant, error)
	Delete(merchant *Merchant) (*Merchant, error)
}

type merchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) *merchantRepository {
	return &merchantRepository{db}
}

func (r *merchantRepository) FindAll() ([]Merchant, error) {
	var merchants []Merchant
	if err := r.db.Preload(clause.Associations).Find(&merchants).Error; err != nil {
		return merchants, err
	}
	return merchants, nil
}

func (r *merchantRepository) FindOne(id string) (*Merchant, error) {
	var merchant Merchant
	if err := r.db.Preload(clause.Associations).First(&merchant, id).Error; err != nil {
		return &merchant, err
	}
	return &merchant, nil
}

func (r *merchantRepository) Create(merchant *Merchant) (*Merchant, error) {
	if err := r.db.Create(&merchant).Error; err != nil {
		return merchant, err
	}
	r.db.Preload(clause.Associations).First(&merchant, merchant.ID)
	return merchant, nil
}

func (r *merchantRepository) Update(merchant *Merchant, merchantInput *UpdateMerchantInput) (*Merchant, error) {
	if err := r.db.Model(&merchant).Updates(merchantInput).Error; err != nil {
		return merchant, err
	}
	return merchant, nil
}

func (r *merchantRepository) Delete(merchant *Merchant) (*Merchant, error) {
	if err := r.db.Delete(merchant).Error; err != nil {
		return merchant, err
	}
	return merchant, nil
}
