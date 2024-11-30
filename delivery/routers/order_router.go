package routers

import (
	"inventory/config"
	"inventory/delivery/controllers"
	"inventory/infrastracture"
	"inventory/repository"
	"inventory/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewOrderRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database)  {

	OrderRepository := repository.NewOrderRepository(DB, config.OrderCollection)
	OrderUsecase := usecase.NewOrderUseCase(OrderRepository)
	OrderController := controllers.NewOrderController(OrderUsecase)

	OrderRouter := route.Group("/orders")
	OrderRouter.Use(infrastracture.AuthMiddleware())
	{
		OrderRouter.POST("/", OrderController.CreateOrders)
		OrderRouter.GET("/", OrderController.GetAllOrders)
		OrderRouter.PUT("/:id", OrderController.UpdateOrders)
		OrderRouter.DELETE("/:id", OrderController.DeleteOrders)
		OrderRouter.GET("/:id", OrderController.GetOrdersByID)
	}
}