package handlers

import (
	. "go-project/entities"
	. "go-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type merchantHandler struct {
	service MerchantService
}

func NewMerchantHandler(s MerchantService) *merchantHandler {
	return &merchantHandler{s}
}

func (h *merchantHandler) FindAll(c *gin.Context) {
	merchants, err := h.service.FindAll()
	if err != nil {
		util.APIResponse(c, "Merchant data does not exists", http.StatusNotFound, http.MethodGet, err.Error())
		return
	}
	util.APIResponse(c, "Merchant data retrieved successfully", http.StatusOK, http.MethodGet, merchants)
}

func (h *merchantHandler) FindOne(c *gin.Context) {
	Merchant, err := h.service.FindOne(c.Param("id"))
	if err != nil {
		util.APIResponse(c, "Merchant data not found", http.StatusNotFound, http.MethodGet, err.Error())
		return
	}
	util.APIResponse(c, "Merchant data retrieved successfully", http.StatusOK, http.MethodGet, Merchant)
}

func (h *merchantHandler) Create(c *gin.Context) {
	var input CreateMerchantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		util.APIResponse(c, "Invalid request payload", http.StatusBadRequest, http.MethodPost, err.Error())
		return
	}
	Merchant, err := h.service.Create(input)
	if err != nil {
		util.APIResponse(c, "Error", http.StatusBadRequest, http.MethodPost, err)
		return
	}
	util.APIResponse(c, "Merchant successfully created", http.StatusCreated, http.MethodPost, Merchant)
}

func (h *merchantHandler) Update(c *gin.Context) {
	var input UpdateMerchantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		util.APIResponse(c, "Invalid request payload", http.StatusBadRequest, http.MethodGet, err.Error())
		return
	}
	Merchant, err := h.service.Update(c.Param("id"), &input)
	if err != nil {
		util.APIResponse(c, "Error", http.StatusBadRequest, http.MethodPatch, err.Error())
		return
	}
	util.APIResponse(c, "Merchant successfully updated", http.StatusOK, http.MethodPatch, Merchant)
}

func (h *merchantHandler) Delete(c *gin.Context) {
	Merchant, err := h.service.Delete(c.Param("id"))
	if err != nil {
		util.APIResponse(c, "Merchant data not found", http.StatusNotFound, http.MethodGet, err.Error())
		return
	}
	util.APIResponse(c, "Merchant data deleted successfully", http.StatusOK, http.MethodDelete, Merchant)
}
