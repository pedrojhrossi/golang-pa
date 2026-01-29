package services

import (
	"context"

	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
	"gitlab.com/pedrojhrossi/golang-pa/internal/ports"
)

type tenantService struct {
	repo ports.TenantRepository
}

func NewTenantService(repo ports.TenantRepository) ports.TenantService {
	return &tenantService{
		repo: repo,
	}
}

func (s *tenantService) RegisterTenant(ctx context.Context, name, email string) (*domain.Tenant, error) {
	newTenant := domain.NewTenant(name, email)

	err := s.repo.Create(ctx, newTenant)
	if err != nil {
		return nil, err
	}

	return newTenant, nil
}
