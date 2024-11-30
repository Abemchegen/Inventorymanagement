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

func NewSupplierRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {
	supplierRepository := repository.NewSupplierRepository(DB, config.SellerCollection)
	supplierUsecase := usecase.NewSupplierUseCase(supplierRepository)
	supplierController := controllers.NewSupplierController(supplierUsecase)

	supplierRoutes := route.Group("/supplier")
	supplierRoutes.Use(infrastracture.AuthMiddleware())
	{
		supplierRoutes.POST("/", supplierController.CreateSuppliers)
		supplierRoutes.GET("/", supplierController.GetAllSuppliers)
		supplierRoutes.GET("/:id", supplierController.GetSuppliersByID)
		supplierRoutes.DELETE("/:id", supplierController.DeleteSuppliers)
		supplierRoutes.PUT("/:id", supplierController.UpdateSuppliers)

	}
}