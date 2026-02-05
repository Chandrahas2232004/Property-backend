package services

import (
	"context"

	"property-backend/models"
	"property-backend/repositories"
)

// PropertyService defines property domain logic
type PropertyService interface {
	Total(ctx context.Context) (int64, error)
	ActiveRentalCount(ctx context.Context) (int64, error)
	AddProperty(ctx context.Context, req interface{}) (int64, error)
	ListByType(ctx context.Context, propertyType string) ([]models.Property, error)
}

type propertyService struct {
	repo repositories.PropertyRepository
}

// NewPropertyService constructs a PropertyService
func NewPropertyService(repo repositories.PropertyRepository) PropertyService {
	return &propertyService{repo: repo}
}

func (s *propertyService) Total(ctx context.Context) (int64, error) {
	return s.repo.Total(ctx)
}

func (s *propertyService) ActiveRentalCount(ctx context.Context) (int64, error) {
	return s.repo.ActiveRentalCount(ctx)
}

func (s *propertyService) AddProperty(ctx context.Context, req interface{}) (int64, error) {
	return s.repo.Create(ctx, req)
}

func (s *propertyService) ListByType(ctx context.Context, propertyType string) ([]models.Property, error) {
	return s.repo.ListByType(ctx, propertyType)
}
