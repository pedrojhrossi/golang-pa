package repository

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
)

// ** tenantDTO defines how an tenant is stored in MongoDB
type tenantDTO struct {
	ID        string     `bson:"_id"`
	Name      string     `bson:"name"`
	Email     string     `bson:"email"`
	Status    string     `bson:"status"`
	CreatedAt time.Time  `bson:"created_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}

func toTenantPersistenceDTO(tenant *domain.Tenant) tenantDTO {
	dto := tenantDTO{
		ID:        tenant.ID.String(),
		Name:      tenant.Name,
		Email:     tenant.Email,
		Status:    tenant.Status,
		CreatedAt: tenant.CreatedAt,
	}

	if tenant.Status == domain.TenantStatusArchived {
		now := time.Now()
		dto.DeletedAt = &now
	}

	return dto
}

func toTenantDomainEntity(dto tenantDTO) *domain.Tenant {
	return &domain.Tenant{
		ID:        uuid.MustParse(dto.ID),
		Name:      dto.Name,
		Email:     dto.Email,
		Status:    dto.Status,
		CreatedAt: dto.CreatedAt,
	}
}
