package controllers

import (
	"go-project/db"
	"go-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func FindAllOrders(c *gin.Context) {
	var orders []models.Order
	db.DB.Preload(clause.Associations).Preload("OrderItems." + clause.Associations).Find(&orders)
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func FindOneOrder(c *gin.Context) {
	var order models.Order

	if err := db.DB.Preload(clause.Associations).Preload("OrderItems."+clause.Associations).First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func CreateOrder(c *gin.Context) {
	var input models.CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		Status: input.Status,
		UserID: input.UserID,
	}

	if err := db.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		var orderItems []models.OrderItem
		for _, itemInput := range input.CreateOrderItemInputs {
			orderItem := models.OrderItem{OrderID: order.ID, Quantity: itemInput.Quantity, ProductID: itemInput.ProductID}
			orderItems = append(orderItems, orderItem)
		}

		if err := tx.Create(&orderItems).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Preload(clause.Associations).Preload("OrderItems." + clause.Associations).Find(&order)

	c.JSON(http.StatusCreated, gin.H{"data": order})
}

// func UpdateProduct(c *gin.Context) {
// 	var product models.Product
// 	if err := db.DB.First(&product, c.Param("id")).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	var input models.UpdateProductInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := db.DB.Model(&product).Updates(input).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": product})
// }

// func DeleteProduct(c *gin.Context) {
// 	var product models.Product
// 	if err := db.DB.First(&product, c.Param("id")).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	if err := db.DB.Delete(&product).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": product})
// }
