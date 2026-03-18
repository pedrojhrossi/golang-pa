package services

import (
	"context"
	"errors"

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

func (s *tenantService) GetTenant(ctx context.Context, id string) (*domain.Tenant, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	return s.repo.GetByID(ctx, id)
}

func (s *tenantService) ListAllTenants(ctx context.Context, includeArchived bool) ([]*domain.Tenant, error) {
	return s.repo.ListTenants(ctx, includeArchived)
}
