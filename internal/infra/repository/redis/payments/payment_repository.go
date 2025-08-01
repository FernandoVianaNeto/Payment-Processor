package redis_payment_repository

import (
	"context"
	domain_repository "payment-gateway/internal/domain/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentsRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewPaymentsRepository(db *mongo.Database) domain_repository.PaymentRepositoryInterface {
	// collection := db.Collection(configs.MongoCfg.WorkoutCollection)
	collection := db.Collection("payments")

	return &PaymentsRepository{
		db:         db,
		collection: collection,
	}
}

func (f *PaymentsRepository) Create(ctx context.Context) error {
	return nil
}

func (f *PaymentsRepository) Summary(ctx context.Context) error {
	return nil
}
