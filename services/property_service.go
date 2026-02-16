package services

import (
	"context"

	"property-backend/models"
	"property-backend/repositories"
)

// PropertyService defines property domain logic
type PropertyService interface {
	AddPropertyBasicInfo(ctx context.Context, req interface{}) (int64, error)
	UploadPropertyFiles(ctx context.Context, filesData interface{}) (map[string]string, error)
	GetAll(ctx context.Context) ([]models.Property, error)
	GetAllByUser(ctx context.Context, userID uint) ([]models.Property, error)
	ListByType(ctx context.Context, propertyType string) ([]models.Property, error)
	ListByTypeByUser(ctx context.Context, propertyType string, userID uint) ([]models.Property, error)
}

type propertyService struct {
	repo repositories.PropertyRepository
}

// NewPropertyService constructs a PropertyService
func NewPropertyService(repo repositories.PropertyRepository) PropertyService {
	return &propertyService{repo: repo}
}

func (s *propertyService) AddPropertyBasicInfo(ctx context.Context, req interface{}) (int64, error) {
	return s.repo.CreateBasicInfo(ctx, req)
}

func (s *propertyService) UploadPropertyFiles(ctx context.Context, filesData interface{}) (map[string]string, error) {
	return s.repo.UploadFiles(ctx, filesData)
}

func (s *propertyService) GetAll(ctx context.Context) ([]models.Property, error) {
	return s.repo.ListAll(ctx)
}

func (s *propertyService) GetAllByUser(ctx context.Context, userID uint) ([]models.Property, error) {
	return s.repo.ListAllByUser(ctx, userID)
}

func (s *propertyService) ListByType(ctx context.Context, propertyType string) ([]models.Property, error) {
	return s.repo.ListByType(ctx, propertyType)
}

func (s *propertyService) ListByTypeByUser(ctx context.Context, propertyType string, userID uint) ([]models.Property, error) {
	return s.repo.ListByTypeByUser(ctx, propertyType, userID)
}
