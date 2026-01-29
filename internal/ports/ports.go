package ports

import (
	"context"

	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *domain.Tenant) error
	GetByID(ctx context.Context, id string) (*domain.Tenant, error)
}

type TenantService interface {
	RegisterTenant(ctx context.Context, name, email string) (*domain.Tenant, error)
}
