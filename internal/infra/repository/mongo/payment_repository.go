package mongo_infra

import (
	"context"
	"errors"
	configs "payment-gateway/cmd/config"
	"payment-gateway/internal/domain/dto"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_response "payment-gateway/internal/domain/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewPaymentRepository(db *mongo.Database) domain_repository.PaymentRepositoryInterface {
	collection := db.Collection(configs.MongoCfg.PaymentCollection)

	return &PaymentRepository{
		db:         db,
		collection: collection,
	}
}

func (f *PaymentRepository) Create(ctx context.Context, input dto.CreatePaymentDto) error {
	_, err := f.collection.InsertOne(ctx, PaymentModel{
		CorrelationId:        input.CorrelationId,
		RequestedAt:          input.RequestedAt,
		Amount:               input.Amount,
		TransactionProcessor: input.TransactionProcessor,
	})

	return err
}

func (f *PaymentRepository) AlreadyAdded(ctx context.Context, correlationId string) bool {
	var model PaymentModel

	filter := bson.M{
		"correlationId": correlationId,
	}

	err := f.collection.FindOne(ctx, filter).Decode(&model)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false
		}

		return false
	}

	return true
}

func (f *PaymentRepository) Summary(ctx context.Context, input dto.GetPaymentsSummaryDto) (*domain_response.PaymentSummaryResponse, error) {
	return nil, nil
}
