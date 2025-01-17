package controllers

import (
	"net/http"
	"pdfGenerator/app/usecases"
	healty_repository "pdfGenerator/infrastructure/model/healthyCheck/repository"

	"github.com/gin-gonic/gin"
)

var (
	heathy_repository = healty_repository.HealthyRepository{}
)

func HealthCheck(c *gin.Context) {
	a := usecases.HealthCheckUseCase(heathy_repository)
	c.JSON(http.StatusOK, gin.H{"Message": a})
}
