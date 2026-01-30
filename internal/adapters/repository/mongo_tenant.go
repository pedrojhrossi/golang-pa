package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
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
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid uuid format")
	}

	var tenant domain.Tenant
	err = r.db.FindOne(ctx, bson.M{"_id": parsedId}).Decode(&tenant)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("tenant not found")
		}
		return nil, err
	}

	return &tenant, nil
}
