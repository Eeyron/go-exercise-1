package handlers

import (
	. "go-project/entities"
	. "go-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/restuwahyu13/gin-rest-api/utils"
)

type userHandler struct {
	service UserService
}

func NewUserHandler(s UserService) *userHandler {
	return &userHandler{s}
}

func (h *userHandler) FindAll(c *gin.Context) {
	users, err := h.service.FindAll()
	if err != nil {
		util.APIResponse(c, "User data does not exists", http.StatusNotFound, http.MethodGet, err.Error())
		return
	}
	util.APIResponse(c, "User data retrieved successfully", http.StatusOK, http.MethodGet, users)
}

func (h *userHandler) FindOne(c *gin.Context) {
	user, err := h.service.FindOne(c.Param("id"))
	if err != nil {
		util.APIResponse(c, "User data not found", http.StatusNotFound, http.MethodGet, err.Error())
		return
	}
	util.APIResponse(c, "User data retrieved successfully", http.StatusOK, http.MethodGet, user)
}

func (h *userHandler) Create(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		util.APIResponse(c, "Invalid request payload", http.StatusBadRequest, http.MethodPost, err.Error())
		return
	}
	user, err := h.service.Create(input)
	if err != nil {
		util.APIResponse(c, "Error", http.StatusBadRequest, http.MethodPost, err)
		return
	}
	util.APIResponse(c, "User successfully created", http.StatusCreated, http.MethodPost, user)
}

func (h *userHandler) Update(c *gin.Context) {
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		util.APIResponse(c, "Invalid request payload", http.StatusBadRequest, http.MethodGet, err.Error())
		return
	}
	user, err := h.service.Update(c.Param("id"), &input)
	if err != nil {
		util.APIResponse(c, "Error", http.StatusBadRequest, http.MethodPatch, err.Error())
		return
	}
	util.APIResponse(c, "User successfully updated", http.StatusOK, http.MethodPatch, user)
}

func (h *userHandler) Delete(c *gin.Context) {
	user, err := h.service.Delete(c.Param("id"))
	if err != nil {
		util.APIResponse(c, "User data not found", http.StatusNotFound, http.MethodGet, err.Error())
		return
	}
	util.APIResponse(c, "User data deleted successfully", http.StatusOK, http.MethodDelete, user)
}
