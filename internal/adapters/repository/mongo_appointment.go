package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gitlab.com/pedrojhrossi/golang-pa/internal/core/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoAppointmentRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewMongoAppointmentRepository(db *mongo.Database) *MongoAppointmentRepository {
	return &MongoAppointmentRepository{
		db:         db,
		collection: db.Collection("appointments"),
	}
}

func (r *MongoAppointmentRepository) Create(ctx context.Context, apt *domain.Appointment) error {
	dto := toPersistenceDTO(apt)

	_, err := r.collection.InsertOne(ctx, dto)
	return err
}

func (r *MongoAppointmentRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Appointment, error) {
	var dto appointmentDTO

	filter := bson.M{
		"_id": id.String(),
	}

	err := r.collection.FindOne(ctx, filter).Decode(&dto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("appointment not found")
		}
		return nil, err
	}

	return toDomainEntity(dto), nil
}

func (r *MongoAppointmentRepository) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*domain.Appointment, error) {
	filter := bson.M{
		"tenant_id": tenantID.String(),
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.New("failed to fetch appointments")
	}
	defer cursor.Close(ctx)

	var dtos []appointmentDTO
	if err := cursor.All(ctx, &dtos); err != nil {
		return nil, err
	}

	appointments := make([]*domain.Appointment, len(dtos))
	for i, d := range dtos {
		appointments[i] = toDomainEntity(d)
	}

	return appointments, nil
}

func (r *MongoAppointmentRepository) FindOverlapping(ctx context.Context, tenantID uuid.UUID, start, end time.Time) ([]*domain.Appointment, error) {
	//** Query logic: (ExistingStart < NewEnd) AND (ExistingEnd > NewStart)
	filter := bson.M{
		"tenant_id":  tenantID.String(),
		"status":     string(domain.StatusScheduled),
		"start_time": bson.M{"$lte": end},
		"end_time":   bson.M{"$gt": start},
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var dtos []appointmentDTO
	if err := cursor.All(ctx, &dtos); err != nil {
		return nil, err
	}

	//** Map DTOs back to Domain entities
	appointments := make([]*domain.Appointment, len(dtos))
	for i, d := range dtos {
		appointments[i] = toDomainEntity(d)
	}

	return appointments, nil
}

func (r *MongoAppointmentRepository) Update(ctx context.Context, apt *domain.Appointment) error {
	dto := toPersistenceDTO(apt)

	filter := bson.M{"_id": dto.ID}
	update := bson.M{"$set": bson.M{"status": dto.Status}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err
}
