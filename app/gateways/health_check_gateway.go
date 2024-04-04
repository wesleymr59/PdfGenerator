package gateways

import "pdfGenerator/app/entities"

type HealthCheckGateway interface {
	HealthCheck() entities.PayloadHealthyCheck
}
