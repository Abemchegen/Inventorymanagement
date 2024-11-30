package controllers

import (
	"inventory/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	OrderUsecase domain.ReportUseCase
}

func NewReportController(reportUsecase domain.ReportUseCase) *ReportController {
	return &ReportController{
		OrderUsecase: reportUsecase,
	}
}

func (a *ReportController) GetBestSellingProduct(c *gin.Context) {
	reports, err := a.OrderUsecase.GetBestSellingProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Failed to get best selling product", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Best selling product retrieved successfully", "data": reports})
}

func (a *ReportController) GetBestSellingCategory(c *gin.Context) {
	categories, err := a.OrderUsecase.GetBestSellingCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Failed to get best selling category", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Best selling category retrieved successfully", "data": categories})
}

func (a *ReportController) GetOverView(c *gin.Context) {
	overview, err := a.OrderUsecase.GetOverView()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Failed to get overview", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Overview retrieved successfully", "data": overview})
}
