package controllers

import (
	"net/http"
	"pdfGenerator/app/usecases"

	"github.com/gin-gonic/gin"
)

func PdfGenerationController(c *gin.Context) {
	a := usecases.HealthCheckUseCase(heathy_repository)
	c.JSON(http.StatusOK, gin.H{"Message": a})
}
