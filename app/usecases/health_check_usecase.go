package usecases

import (
	"fmt"
	"pdfGenerator/app/gateways"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	once                    sync.Once
	mu                      sync.Mutex
	Health_Check_repository gateways.HealthCheckGateway
)

func HealthCheckUseCase(health_repo gateways.HealthCheckGateway, c *gin.Context) {
	once.Do(func() { initialize(health_repo) })
	fmt.Println("aaaaaaaaaaa")
	fmt.Println(Health_Check_repository.HealthCheck())
}
