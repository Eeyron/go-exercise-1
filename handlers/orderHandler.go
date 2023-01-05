package handlers

import (
	. "go-project/entities"
	. "go-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type orderHandler struct {
	service OrderService
}

func NewOrderHandler(s OrderService) *orderHandler {
	return &orderHandler{s}
}

func (h *orderHandler) FindAll(c *gin.Context) {
	orders, err := h.service.FindAll()
	if err != nil {
		util.APIResponse(c, "Order data does not exists", http.StatusNotFound, http.MethodGet, err.Error())
		return
	}
	util.APIResponse(c, "Order data retrieved successfully", http.StatusOK, http.MethodGet, orders)
}

func (h *orderHandler) FindOne(c *gin.Context) {
	Order, err := h.service.FindOne(c.Param("id"))
	if err != nil {
		util.APIResponse(c, "Order data not found", http.StatusNotFound, http.MethodGet, err.Error())
		return
	}
	util.APIResponse(c, "Order data retrieved successfully", http.StatusOK, http.MethodGet, Order)
}

func (h *orderHandler) Create(c *gin.Context) {
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		util.APIResponse(c, "Invalid request payload", http.StatusBadRequest, http.MethodPost, err.Error())
		return
	}
	Order, err := h.service.Create(input)
	if err != nil {
		util.APIResponse(c, "Error", http.StatusBadRequest, http.MethodPost, err)
		return
	}
	util.APIResponse(c, "Order successfully created", http.StatusCreated, http.MethodPost, Order)
}
