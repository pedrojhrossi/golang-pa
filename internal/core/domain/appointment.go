package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorOverlap     = errors.New("appointment overlaps with an existing schedule")
	ErrorInvalidTime = errors.New("start time must be before end time")
	ErrorNotFound    = errors.New("appointment not found")
	ErrorInvalidID   = errors.New("invalid or empty appointment id")
)

type AppointmentStatus string

const (
	StatusScheduled AppointmentStatus = "scheduled"
	StatusCancelled AppointmentStatus = "cancelled"
	StatusCompleted AppointmentStatus = "completed"
)

type Appointment struct {
	ID          uuid.UUID
	TenantID    uuid.UUID
	PatientID   uuid.UUID
	PatientName string
	StartTime   time.Time
	EndTime     time.Time
	Status      AppointmentStatus
	CreatedAt   time.Time
}

// ** NewAppointment is a factory function that ensures a valid domain object
func NewAppointment(tenantID, patientID uuid.UUID, patientName string, start, end time.Time) (*Appointment, error) {
	if start.After(end) || start.Equal(end) {
		return nil, ErrorInvalidTime
	}

	return &Appointment{
		ID:          uuid.New(),
		TenantID:    tenantID,
		PatientID:   patientID,
		PatientName: patientName,
		StartTime:   start,
		EndTime:     end,
		Status:      StatusScheduled,
		CreatedAt:   time.Now(),
	}, nil
}

func (a *Appointment) Cancel() error {
	if a.Status == StatusCancelled {
		return errors.New("appointment is already cancelled")
	}

	if a.StartTime.Before(time.Now()) {
		return errors.New("cannot cancel past appointments")
	}

	a.Status = StatusCancelled
	return nil
}
