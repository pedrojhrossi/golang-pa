package mocks

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
)

type AppointmentRepository struct {
	mock.Mock
}

func (m *AppointmentRepository) Create(ctx context.Context, appointment *domain.Appointment) error {
	args := m.Called(ctx, appointment)
	return args.Error(0)
}

func (m *AppointmentRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Appointment, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Appointment), args.Error(1)
}

func (m *AppointmentRepository) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*domain.Appointment, error) {
	args := m.Called(ctx, tenantID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Appointment), args.Error(1)
}

func (m *AppointmentRepository) FindOverlapping(ctx context.Context, tenantID uuid.UUID, start, end time.Time) ([]*domain.Appointment, error) {
	args := m.Called(ctx, tenantID, start, end)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Appointment), args.Error(1)
}

func (m *AppointmentRepository) Update(ctx context.Context, appointment *domain.Appointment) error {
	args := m.Called(ctx, appointment)
	return args.Error(0)
}
