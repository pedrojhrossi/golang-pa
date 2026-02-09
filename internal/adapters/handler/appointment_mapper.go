package handler

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
)

// ** scheduleRequest is the DTO fot incoming JSON
type scheduleRequest struct {
	PatientID   string    `json:"patient_id"`
	PatientName string    `json:"patient_name"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

// ** appointmentResponse is the DTO for outgoing JSON
type appointmentResponse struct {
	ID          string    `json:"id"`
	PatientName string    `json:"patient_name"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Status      string    `json:"status"`
}

// ** toDomain converts the Request DTO to Domain variables
func (r *scheduleRequest) toDomain() (uuid.UUID, error) {
	return uuid.Parse(r.PatientID)
}

// ** fromDomain converts Domain Entity into a Response DTO
func fromDomain(apt *domain.Appointment) appointmentResponse {
	return appointmentResponse{
		ID:          apt.ID.String(),
		PatientName: apt.PatientName,
		StartTime:   apt.StartTime,
		EndTime:     apt.EndTime,
		Status:      string(apt.Status),
	}
}
