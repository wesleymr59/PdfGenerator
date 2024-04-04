package healty_repository

import "pdfGenerator/app/entities"

type HealthyRepository struct{}

func (x HealthyRepository) HealthCheck() entities.PayloadHealthyCheck {
	payload := entities.PayloadHealthyCheck{Message: "deu bom"}
	return payload
}
