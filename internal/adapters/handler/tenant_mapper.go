package handler

import (
	"time"

	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
)

type tenantRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type tenantResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func fromTenantDomain(tenant *domain.Tenant) tenantResponse {
	return tenantResponse{
		ID:        tenant.ID.String(),
		Name:      tenant.Name,
		Email:     tenant.Email,
		Status:    tenant.Status,
		CreatedAt: tenant.CreatedAt,
	}
}
