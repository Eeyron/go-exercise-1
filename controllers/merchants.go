package controllers

import (
	"go-project/db"
	"go-project/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllMerchants(c *gin.Context) {
	var merchants []models.Merchant
	db.DB.Preload("User").Find(&merchants)
	c.JSON(http.StatusOK, gin.H{"data": merchants})
}

func FindOneMerchant(c *gin.Context) {
	var merchant models.Merchant

	if err := db.DB.Preload("User").First(&merchant, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": merchant})
}

func CreateMerchant(c *gin.Context) {
	var input models.CreateMerchantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merchant := models.Merchant{Name: input.Name, SocialLink: input.SocialLink, UserID: input.UserID}
	log.Print("Merchant1", merchant.User)

	if err := db.DB.Create(&merchant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Preload("User").First(&merchant, merchant.ID)

	c.JSON(http.StatusCreated, gin.H{"data": merchant})
}

func UpdateMerchant(c *gin.Context) {
	var merchant models.Merchant
	if err := db.DB.Preload("User").First(&merchant, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateMerchantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Model(&merchant).Updates(input).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": merchant})
}

func DeleteMerchant(c *gin.Context) {
	var merchant models.Merchant
	if err := db.DB.Where("id = ?", c.Param("id")).First(&merchant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.DB.Delete(&merchant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": merchant})
}
