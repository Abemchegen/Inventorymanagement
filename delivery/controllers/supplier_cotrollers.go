package controllers

import (
	"inventory/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SupplierController struct {
	SupplierUsecase domain.SuppliersUseCase
}

func NewSupplierController(SupplierUsecase domain.SuppliersUseCase) *SupplierController {
	return &SupplierController{
		SupplierUsecase: SupplierUsecase,
	}
}

func (s *SupplierController) CreateSuppliers(c *gin.Context) {
	var supplier domain.Suppliers
	c.BindJSON(&supplier)
	result, err := s.SupplierUsecase.CreateSuppliers(supplier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Supplier created successfully", "data": result})
}

func (s *SupplierController) GetAllSuppliers(c *gin.Context) {
	result, err := s.SupplierUsecase.GetAllSuppliers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Suppliers retrieved successfully", "data": result})
}

func (s *SupplierController) GetSuppliersByID(c *gin.Context) {
	id := c.Param("id")
	result, err := s.SupplierUsecase.GetSuppliersByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Supplier retrieved successfully", "data": result})
}

func (s *SupplierController) DeleteSuppliers(c *gin.Context) {
	id := c.Param("id")
	_, err := s.SupplierUsecase.DeleteSuppliers(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Supplier deleted successfully"})
}

func (s *SupplierController) UpdateSuppliers(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid ID"})
		return
	}
	
	var supplier domain.Suppliers
	c.BindJSON(&supplier)
	supplier.ID = objectID
	result, err := s.SupplierUsecase.UpdateSuppliers(supplier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Supplier updated successfully", "data": result})
}
