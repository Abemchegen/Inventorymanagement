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

func NewReportRouter(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	reportRepository := repository.NewReportRepository(DB,config.OrderCollection, config.ProductCollection, config.SellerCollection, config.StoreCollection)
	reportUsecase := usecase.NewReportUseCase(reportRepository)
	reportController := controllers.NewReportController(reportUsecase)

	reportRouter := route.Group("/report")
	reportRouter.Use(infrastracture.AuthMiddleware())
	{
		reportRouter.GET("/best-selling-product", reportController.GetBestSellingProduct)
		reportRouter.GET("/best-selling-category", reportController.GetBestSellingCategory)
		reportRouter.GET("/overview", reportController.GetOverView)
	}

}