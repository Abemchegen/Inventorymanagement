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

func NewStoreRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	StoreRepository := repository.NewStoreRepository(DB, config.StoreCollection)
	StoreUseCase := usecase.NewStoreUseCase(StoreRepository)
	StoreController := controllers.NewStoreController(StoreUseCase)

	StoreRouter := route.Group("/store")
	StoreRouter.Use(infrastracture.AuthMiddleware())
	{
		StoreRouter.POST("/", StoreController.CreateStore)
		StoreRouter.GET("/", StoreController.GetAllStore)
		StoreRouter.GET("/:id", StoreController.GetStoreByID)
		StoreRouter.PUT("/:id", StoreController.UpdateStore)
		StoreRouter.DELETE("/:id", StoreController.DeleteStore)
	}
	


}