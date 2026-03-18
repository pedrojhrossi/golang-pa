package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	TenantStatusActive   = "active"
	TenantStatusArchived = "archived"
)

type Tenant struct {
	ID        uuid.UUID `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Status    string    `bson:"status"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func NewTenant(name, email string) *Tenant {
	return &Tenant{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Status:    TenantStatusActive,
		CreatedAt: time.Now(),
	}
}

func (t *Tenant) Archive() {
	t.Status = TenantStatusArchived
}

func (t *Tenant) IsActive() bool {
	return t.Status == TenantStatusActive
}
