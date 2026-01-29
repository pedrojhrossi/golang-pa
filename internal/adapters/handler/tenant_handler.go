package handler

import (
	"encoding/json"
	"net/http"

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
	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

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
	json.NewEncoder(w).Encode(tenant)
}
