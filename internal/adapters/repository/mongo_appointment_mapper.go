package repository

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
)

// ** appointmentDTO defines how an appointment is stored in MongoDB
type appointmentDTO struct {
	ID          string    `bson:"_id"`
	TenantID    string    `bson:"tenant_id"`
	PatientID   string    `bson:"patient_id"`
	PatientName string    `bson:"patient_name"`
	StartTime   time.Time `bson:"start_time"`
	EndTime     time.Time `bson:"end_time"`
	Status      string    `bson:"status"`
	CreatedAt   time.Time `bson:"created_at"`
}

// ** toPersistenceDTO maps from Domain -> DTO (for saving)
func toPersistenceDTO(apt *domain.Appointment) appointmentDTO {
	return appointmentDTO{
		ID:          apt.ID.String(),
		TenantID:    apt.TenantID.String(),
		PatientID:   apt.PatientID.String(),
		PatientName: apt.PatientName,
		StartTime:   apt.StartTime,
		EndTime:     apt.EndTime,
		Status:      string(apt.Status),
		CreatedAt:   apt.CreatedAt,
	}
}

// ** toDomainEntity maps from DTO -> Domain (for reading)
func toDomainEntity(dto appointmentDTO) *domain.Appointment {
	return &domain.Appointment{
		ID:          uuid.MustParse(dto.ID),
		TenantID:    uuid.MustParse(dto.TenantID),
		PatientID:   uuid.MustParse(dto.PatientID),
		PatientName: dto.PatientName,
		StartTime:   dto.StartTime,
		EndTime:     dto.EndTime,
		Status:      domain.AppointmentStatus(dto.Status),
		CreatedAt:   dto.CreatedAt,
	}
}
