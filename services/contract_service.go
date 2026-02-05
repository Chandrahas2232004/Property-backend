package services

import (
	"context"

	"property-backend/models"
	"property-backend/repositories"
)

// ContractService defines contract domain logic
type ContractService interface {
	AddContract(ctx context.Context, c *models.Contract) (int64, error)
	GetAllContracts(ctx context.Context) ([]models.Contract, error)
	GetContractsByType(ctx context.Context, contractType string) ([]models.Contract, error)
}

type contractService struct {
	repo repositories.ContractRepository
}

// NewContractService constructs a ContractService
func NewContractService(repo repositories.ContractRepository) ContractService {
	return &contractService{repo: repo}
}

func (s *contractService) AddContract(ctx context.Context, c *models.Contract) (int64, error) {
	return s.repo.Create(ctx, c)
}

func (s *contractService) GetAllContracts(ctx context.Context) ([]models.Contract, error) {
	return s.repo.ListAll(ctx)
}

func (s *contractService) GetContractsByType(ctx context.Context, contractType string) ([]models.Contract, error) {
	return s.repo.ListByType(ctx, contractType)
}
