package services

import (
	"context"

	"property-backend/repositories"
)

// FormService defines form domain logic
type FormService interface {
	GetCountry(ctx context.Context) ([]repositories.FormData, error)
	GetState(ctx context.Context, countryID string) ([]repositories.FormData, error)
	GetDistrict(ctx context.Context, stateID string) ([]repositories.FormData, error)
	GetTaluk(ctx context.Context, districtID string) ([]repositories.FormData, error)
}

type formService struct {
	repo repositories.FormRepository
}

// NewFormService constructs a FormService
func NewFormService(repo repositories.FormRepository) FormService {
	return &formService{repo: repo}
}

func (s *formService) GetCountry(ctx context.Context) ([]repositories.FormData, error) {
	return s.repo.GetCountry(ctx)
}

func (s *formService) GetState(ctx context.Context, countryID string) ([]repositories.FormData, error) {
	return s.repo.GetState(ctx, countryID)
}

func (s *formService) GetDistrict(ctx context.Context, stateID string) ([]repositories.FormData, error) {
	return s.repo.GetDistrict(ctx, stateID)
}

func (s *formService) GetTaluk(ctx context.Context, districtID string) ([]repositories.FormData, error) {
	return s.repo.GetTaluk(ctx, districtID)
}
