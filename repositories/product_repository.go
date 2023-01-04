package repositories

import (
	. "go-project/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	FindAll() ([]Product, error)
	FindOne(id string) (*Product, error)
	Create(product *Product) (*Product, error)
	Update(product *Product, productInput *UpdateProductInput) (*Product, error)
	Delete(product *Product) (*Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll() ([]Product, error) {
	var products []Product
	if err := r.db.Preload(clause.Associations).Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func (r *productRepository) FindOne(id string) (*Product, error) {
	var product Product
	if err := r.db.Preload(clause.Associations).First(&product, id).Error; err != nil {
		return &product, err
	}
	return &product, nil
}

func (r *productRepository) Create(product *Product) (*Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return product, err
	}
	r.db.Preload(clause.Associations).First(&product, product.ID)
	return product, nil
}

func (r *productRepository) Update(product *Product, productInput *UpdateProductInput) (*Product, error) {
	if err := r.db.Model(&product).Updates(productInput).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *productRepository) Delete(product *Product) (*Product, error) {
	if err := r.db.Delete(product).Error; err != nil {
		return product, err
	}
	return product, nil
}
