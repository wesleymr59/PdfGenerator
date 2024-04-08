package usecases

import (
	"fmt"
	"pdfGenerator/app/entities"
	"pdfGenerator/app/gateways"
	"sync"
)

var (
	once                    sync.Once
	mu                      sync.Mutex
	Health_Check_repository gateways.HealthCheckGateway
)

func HealthCheckUseCase(health_repo gateways.HealthCheckGateway) entities.PayloadHealthyCheck {
	fmt.Println(health_repo.HealthCheck())
	return health_repo.HealthCheck()
}
