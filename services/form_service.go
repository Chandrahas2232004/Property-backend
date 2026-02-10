package services

import (
	"context"

	"property-backend/repositories"
)

// FormService defines form domain logic
type FormService interface {
	GetCountry(ctx context.Context) ([]string, error)
	GetState(ctx context.Context, countryID string) ([]string, error)
	GetDistrict(ctx context.Context, stateID string) ([]string, error)
	GetTaluk(ctx context.Context, districtID string) ([]string, error)
}

type formService struct {
	repo repositories.FormRepository
}

// NewFormService constructs a FormService
func NewFormService(repo repositories.FormRepository) FormService {
	return &formService{repo: repo}
}

func (s *formService) GetCountry(ctx context.Context) ([]string, error) {
	return s.repo.GetCountry(ctx)
}

func (s *formService) GetState(ctx context.Context, countryID string) ([]string, error) {
	return s.repo.GetState(ctx, countryID)
}

func (s *formService) GetDistrict(ctx context.Context, stateID string) ([]string, error) {
	return s.repo.GetDistrict(ctx, stateID)
}

func (s *formService) GetTaluk(ctx context.Context, districtID string) ([]string, error) {
	return s.repo.GetTaluk(ctx, districtID)
}
