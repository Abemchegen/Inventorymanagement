package routers

import (
	"inventory/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(route *gin.RouterGroup, config *config.Config, DB mongo.Database) {

	NewUserRouter(route , config, DB)
	NewProductRouter(route , config, DB)
	NewOrderRouter(route , config, DB)
	NewSupplierRouter(route , config, DB)
	NewReportRouter(route , config, DB)
	NewStoreRouter(route , config, DB)

}