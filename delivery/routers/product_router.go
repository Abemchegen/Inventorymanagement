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

func NewProductRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	productRepository := repository.NewProductRepository(DB, config.ProductCollection)
	productUsecase := usecase.NewProductUseCase(productRepository)
	productController := controllers.NewProductController(productUsecase)

	productRouter :=  route.Group("/product")
	productRouter.Use(infrastracture.AuthMiddleware())
	{
		productRouter.POST("/", productController.CreateProduct)
		productRouter.GET("/", productController.GetAllProduct)
		productRouter.GET("/:id", productController.GetProductByID)
		productRouter.PUT("/:id", productController.UpdateProduct)
		productRouter.DELETE("/:id", productController.DeleteProduct)

	}
	
}
