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
	db         *mongo.Database
	collection *mongo.Collection
}

func NewMongoTenantRepository(db *mongo.Database) *MongoTenantRepository {
	return &MongoTenantRepository{
		db:         db,
		collection: db.Collection("tenants"),
	}
}

func (r *MongoTenantRepository) Create(ctx context.Context, tenant *domain.Tenant) error {
	dto := toTenantPersistenceDTO(tenant)
	_, err := r.collection.InsertOne(ctx, dto)
	return err
}

func (r *MongoTenantRepository) GetByID(ctx context.Context, id string) (*domain.Tenant, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid uuid format")
	}

	var dto tenantDTO

	filter := bson.M{
		"_id":    id,
		"status": domain.TenantStatusActive,
	}

	err = r.collection.FindOne(ctx, filter).Decode(&dto)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("tenant not found")
		}
		return nil, err
	}

	return toTenantDomainEntity(dto), nil
}

func (r *MongoTenantRepository) ListTenants(ctx context.Context, includeArchived bool) ([]*domain.Tenant, error) {
	filter := bson.M{"status": domain.TenantStatusActive}

	if includeArchived {
		filter = bson.M{}
	}

	cursor, err := r.collection.Find(ctx, filter)

	if err != nil {
		return nil, errors.New("failed to fetch tenants")
	}
	defer cursor.Close(ctx)

	var dtos []tenantDTO
	if err := cursor.All(ctx, &dtos); err != nil {
		return nil, err
	}

	tenants := make([]*domain.Tenant, len(dtos))
	for i, d := range dtos {
		tenants[i] = toTenantDomainEntity(d)
	}

	return tenants, nil
}
