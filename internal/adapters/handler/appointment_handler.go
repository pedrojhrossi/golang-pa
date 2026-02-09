package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gitlab.com/pedrojhrossi/golang-pa/internal/ports"
)

type AppointmentHandler struct {
	service ports.AppointmentService
}

func NewAppointmentHandler(service ports.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service: service}
}

func (h *AppointmentHandler) Schedule(w http.ResponseWriter, r *http.Request) {
	//** Extract TenantID from URL
	tenantID, err := uuid.Parse(chi.URLParam(r, "tenantID"))
	if err != nil {
		http.Error(w, "Invalid tenant id", http.StatusBadRequest)
		return
	}

	//** Decode Json body
	var req scheduleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	patientID, err := req.toDomain()
	if err != nil {
		http.Error(w, "invalid patient id", http.StatusBadRequest)
		return
	}

	apt, err := h.service.Schedule(r.Context(), tenantID, patientID, req.PatientName, req.StartTime, req.EndTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := fromDomain(apt)
	json.NewEncoder(w).Encode(response)
}

func (h *AppointmentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	apt, err := h.service.GetAppointment(r.Context(), id)
	if err != nil || apt == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := fromDomain(apt)
	json.NewEncoder(w).Encode(response)
}

func (h *AppointmentHandler) List(w http.ResponseWriter, r *http.Request) {
	tenantID, err := uuid.Parse(chi.URLParam(r, "tenantID"))
	if err != nil {
		http.Error(w, "invalid tenant id", http.StatusBadRequest)
		return
	}

	appointments, err := h.service.ListByTenant(r.Context(), tenantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]appointmentResponse, len(appointments))
	for i, a := range appointments {
		response[i] = fromDomain(a)
	}

	json.NewEncoder(w).Encode(response)
}

func (h *AppointmentHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.service.Cancel(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
