package controllers

import (
	"inventory/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderController struct {
	OrderUsecase domain.OrdersUseCase
}

func NewOrderController(OrderUsecase domain.OrdersUseCase) *OrderController {
	return &OrderController{
		OrderUsecase: OrderUsecase,
	}
}

func (a *OrderController) CreateOrders(c *gin.Context) {
	var orders domain.Orders
	err := c.BindJSON(&orders)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid request", "error": err.Error()})
		return
	}
	result, err := a.OrderUsecase.CreateOrders(orders)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to create orders", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Orders created successfully", "data": result})
}

func (a *OrderController) GetAllOrders(c *gin.Context) {
	result, err := a.OrderUsecase.GetAllOrders()
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to get orders", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Orders retrieved successfully", "data": result})
}

func (a *OrderController) UpdateOrders(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid request", "error": "id is required"})
		return
	}
	var orders domain.Orders
	err := c.BindJSON(&orders)

	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid request", "error": err.Error()})
		return
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Invalid ID format", "error": err.Error()})
		return
	}
	orders.ID = objectID
	result, err := a.OrderUsecase.UpdateOrders(orders)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to update orders", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Orders updated successfully", "data": result})
}

func (a *OrderController) DeleteOrders(c *gin.Context) {
	id := c.Param("id")
	result, err := a.OrderUsecase.DeleteOrders(id)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to delete orders", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Orders deleted successfully", "data": result})
}

func (a *OrderController) GetOrdersByID(c *gin.Context) {
	id := c.Param("id")
	result, err := a.OrderUsecase.GetOrdersByID(id)
	if err != nil {
		c.JSON(400, gin.H{"status": 400, "message": "Failed to get orders", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 200, "message": "Orders retrieved successfully", "data": result})
}
