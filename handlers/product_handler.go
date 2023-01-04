package handlers

import (
	. "go-project/entities"
	. "go-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type productHandler struct {
	service ProductService
}

func NewProductHandler(s ProductService) *productHandler {
	return &productHandler{s}
}

func (h *productHandler) FindAll(c *gin.Context) {
	products, err := h.service.FindAll()
	if err != nil {
		util.APIResponse(c, "Product data does not exists", http.StatusNotFound, http.MethodGet, err.Error())
		return
	}
	util.APIResponse(c, "Product data retrieved successfully", http.StatusOK, http.MethodGet, products)
}

func (h *productHandler) FindOne(c *gin.Context) {
	Product, err := h.service.FindOne(c.Param("id"))
	if err != nil {
		util.APIResponse(c, "Product data not found", http.StatusNotFound, http.MethodGet, err.Error())
		return
	}
	util.APIResponse(c, "Product data retrieved successfully", http.StatusOK, http.MethodGet, Product)
}

func (h *productHandler) Create(c *gin.Context) {
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		util.APIResponse(c, "Invalid request payload", http.StatusBadRequest, http.MethodPost, err.Error())
		return
	}
	Product, err := h.service.Create(input)
	if err != nil {
		util.APIResponse(c, "Error", http.StatusBadRequest, http.MethodPost, err)
		return
	}
	util.APIResponse(c, "Product successfully created", http.StatusCreated, http.MethodPost, Product)
}

func (h *productHandler) Update(c *gin.Context) {
	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		util.APIResponse(c, "Invalid request payload", http.StatusBadRequest, http.MethodGet, err.Error())
		return
	}
	Product, err := h.service.Update(c.Param("id"), &input)
	if err != nil {
		util.APIResponse(c, "Error", http.StatusBadRequest, http.MethodPatch, err.Error())
		return
	}
	util.APIResponse(c, "Product successfully updated", http.StatusOK, http.MethodPatch, Product)
}

func (h *productHandler) Delete(c *gin.Context) {
	Product, err := h.service.Delete(c.Param("id"))
	if err != nil {
		util.APIResponse(c, "Product data not found", http.StatusNotFound, http.MethodGet, err.Error())
		return
	}
	util.APIResponse(c, "Product data deleted successfully", http.StatusOK, http.MethodDelete, Product)
}
