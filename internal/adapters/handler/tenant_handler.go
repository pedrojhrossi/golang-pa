package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/pedrojhrossi/golang-pa/internal/ports"
)

type TenantHandler struct {
	service ports.TenantService
}

func NewTenantHandler(service ports.TenantService) *TenantHandler {
	return &TenantHandler{
		service: service,
	}
}

func (h *TenantHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body tenantRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	tenant, err := h.service.RegisterTenant(r.Context(), body.Name, body.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := fromTenantDomain(tenant)
	json.NewEncoder(w).Encode(response)
}

func (h *TenantHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	tenant, err := h.service.GetTenant(r.Context(), id)
	if err != nil {
		if err.Error() == "tenant not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := fromTenantDomain(tenant)
	json.NewEncoder(w).Encode(response)
}

func (h *TenantHandler) List(w http.ResponseWriter, r *http.Request) {
	includeArchived := r.URL.Query().Get("archived") == "true"
	tenants, err := h.service.ListAllTenants(r.Context(), includeArchived)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]tenantResponse, len(tenants))
	for i, t := range tenants {
		response[i] = fromTenantDomain(t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TenantHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteTenant(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
