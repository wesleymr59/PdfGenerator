package controllers

import (
	"pdfGenerator/app/usecases"
	healty_repository "pdfGenerator/infrastructure/model/healthyCheck/repository"

	"github.com/gin-gonic/gin"
)

var (
	heathy_repository = healty_repository.HealthyRepository{}
)

func HandleRequests() {
	r := gin.Default()
	// Rotas abertas

	r.POST("/login", usecases.HealthCheckUseCase())
}
