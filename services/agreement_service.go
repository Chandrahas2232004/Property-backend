package services

import (
	"context"

	"property-backend/models"
	"property-backend/repositories"
)

// AgreementService defines agreement domain logic
type AgreementService interface {
	AddRentalAgreement(ctx context.Context, a *models.Agreement) (int64, error)
	GetAllAgreements(ctx context.Context) ([]models.Agreement, error)
}

type agreementService struct {
	repo repositories.AgreementRepository
}

// NewAgreementService constructs an AgreementService
func NewAgreementService(repo repositories.AgreementRepository) AgreementService {
	return &agreementService{repo: repo}
}

func (s *agreementService) AddRentalAgreement(ctx context.Context, a *models.Agreement) (int64, error) {
	return s.repo.CreateRental(ctx, a)
}

func (s *agreementService) GetAllAgreements(ctx context.Context) ([]models.Agreement, error) {
	return s.repo.ListAll(ctx)
}
