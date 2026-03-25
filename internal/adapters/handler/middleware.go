package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
	"gitlab.com/pedrojhrossi/golang-pa/internal/ports"
)

type contextKey string

const TenantContextKey contextKey = "tenant"

func TenantContext(tenantService ports.TenantService) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "tenantID")

			tenant, err := tenantService.GetTenant(r.Context(), id)
			if err != nil {
				http.Error(w, "Resource not found", http.StatusNotFound)
				return
			}

			ctx := context.WithValue(r.Context(), TenantContextKey, tenant)

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetTenantFromContext(ctx context.Context) *domain.Tenant {
	tenant, ok := ctx.Value(TenantContextKey).(*domain.Tenant)
	if !ok {
		return nil
	}

	return tenant
}
