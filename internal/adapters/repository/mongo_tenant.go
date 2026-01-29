package repository

import (
	"context"

	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoTenantRepository struct {
	db *mongo.Collection
}

func NewMongoTenantRepository(client *mongo.Client, dbName string) *MongoTenantRepository {
	return &MongoTenantRepository{
		db: client.Database(dbName).Collection("tenants"),
	}
}

func (r *MongoTenantRepository) Create(ctx context.Context, tenant *domain.Tenant) error {
	_, err := r.db.InsertOne(ctx, tenant)
	return err
}

func (r *MongoTenantRepository) GetByID(ctx context.Context, id string) (*domain.Tenant, error) {
	// Implementation logic for finding by ID
	return nil, nil
}
