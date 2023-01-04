package repositories

import (
	. "go-project/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository interface {
	FindAll() ([]Order, error)
	FindOne(id string) (*Order, error)
	Create(order *Order, itemInputs []CreateOrderItemInput) (*Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) FindAll() ([]Order, error) {
	var orders []Order
	if err := r.db.Preload(clause.Associations).Preload("OrderItems." + clause.Associations).Find(&orders).Error; err != nil {
		return orders, err
	}
	return orders, nil
}

func (r *orderRepository) FindOne(id string) (*Order, error) {
	var order Order
	if err := r.db.Preload(clause.Associations).Preload("OrderItems."+clause.Associations).First(&order, id).Error; err != nil {
		return &order, err
	}
	return &order, nil
}

func (r *orderRepository) Create(order *Order, itemInputs []CreateOrderItemInput) (*Order, error) {
	if err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(order).Error; err != nil {
			return err
		}

		var orderItems []OrderItem
		for _, itemInput := range itemInputs {
			orderItem := OrderItem{OrderID: order.ID, Quantity: itemInput.Quantity, ProductID: itemInput.ProductID}
			orderItems = append(orderItems, orderItem)
		}

		if err := tx.Create(&orderItems).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return order, err
	}

	r.db.Preload(clause.Associations).Preload("OrderItems."+clause.Associations).First(&order, order.ID)
	return order, nil
}
