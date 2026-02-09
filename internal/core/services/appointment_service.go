package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
	"gitlab.com/pedrojhrossi/golang-pa/internal/ports"
)

type appointmentService struct {
	repo ports.AppointmentRepository
}

func NewAppointmentService(repo ports.AppointmentRepository) ports.AppointmentService {
	return &appointmentService{
		repo: repo,
	}
}

func (s *appointmentService) Schedule(ctx context.Context, tenantID, patientID uuid.UUID, name string, start, end time.Time) (*domain.Appointment, error) {
	// TODO:  Logic Check: Does the patient belong to this tenant?
	// Assuming that later a GetByID will be add to a PatientRepository

	existing, err := s.repo.FindOverlapping(ctx, tenantID, start, end)
	if err != nil {
		return nil, err
	}

	if len(existing) > 0 {
		return nil, domain.ErrorOverlap
	}

	apt, err := domain.NewAppointment(tenantID, patientID, name, start, end)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, apt); err != nil {
		return nil, errors.New("Error while creating appointment")
	}

	return apt, nil
}

func (s *appointmentService) GetAppointment(ctx context.Context, id uuid.UUID) (*domain.Appointment, error) {
	if id == uuid.Nil {
		return nil, errors.New("id is required")
	}

	return s.repo.GetByID(ctx, id)
}

func (s *appointmentService) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*domain.Appointment, error) {
	if tenantID == uuid.Nil {
		return nil, errors.New("id is required")
	}

	return s.repo.ListByTenant(ctx, tenantID)
}

func (s *appointmentService) Cancel(ctx context.Context, id uuid.UUID) error {
	//TODO: Implementation to Cancel an appointment pending
	if id == uuid.Nil {
		return domain.ErrorInvalidID
	}

	apt, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if apt == nil {
		return domain.ErrorNotFound
	}

	apt.Cancel()

	return s.repo.Update(ctx, apt)
}
