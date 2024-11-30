package controllers

import (
	"inventory/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StoreController struct {
	StoreUsecase domain.StoreUseCase
}

func NewStoreController(StoreUsecase domain.StoreUseCase) *StoreController {
	return &StoreController{
		StoreUsecase: StoreUsecase,
	}
}

func (s *StoreController) CreateStore(c *gin.Context) {
	var store domain.Store
	if err := c.BindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}
	result, err := s.StoreUsecase.CreateStore(store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create store", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Store created successfully", "data": result})
}

func (s *StoreController) GetAllStore(c *gin.Context) {
	result, err := s.StoreUsecase.GetAllStore()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve stores", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stores retrieved successfully", "data": result})
}

func (s *StoreController) GetStoreByID(c *gin.Context) {
	id := c.Param("id")
	result, err := s.StoreUsecase.GetStoreByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve store", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Store retrieved successfully", "data": result})
}

func (s *StoreController) UpdateStore(c *gin.Context) {
	var store domain.Store
	if err := c.BindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}
	result, err := s.StoreUsecase.UpdateStore(store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update store", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Store updated successfully", "data": result})
}

func (s *StoreController) DeleteStore(c *gin.Context) {
	id := c.Param("id")
	_,err := s.StoreUsecase.DeleteStore(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete store", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Store deleted successfully"})
}
