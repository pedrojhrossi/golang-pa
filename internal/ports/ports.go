package ports

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *domain.Tenant) error
	GetByID(ctx context.Context, id string) (*domain.Tenant, error)
	ListTenants(ctx context.Context) ([]*domain.Tenant, error)
}

type TenantService interface {
	RegisterTenant(ctx context.Context, name, email string) (*domain.Tenant, error)
	GetTenant(ctx context.Context, id string) (*domain.Tenant, error)
	ListAllTenants(ctx context.Context) ([]*domain.Tenant, error)
}

type AppointmentRepository interface {
	Create(ctx context.Context, appointment *domain.Appointment) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Appointment, error)
	ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*domain.Appointment, error)
	FindOverlapping(ctx context.Context, tenantID uuid.UUID, start, end time.Time) ([]*domain.Appointment, error)
	Update(ctx context.Context, apt *domain.Appointment) error
}

type AppointmentService interface {
	Schedule(ctx context.Context, tenantID, patientID uuid.UUID, patientName string, start, end time.Time) (*domain.Appointment, error)
	GetAppointment(ctx context.Context, id uuid.UUID) (*domain.Appointment, error)
	ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*domain.Appointment, error)
	Cancel(ctx context.Context, id uuid.UUID) error
}
