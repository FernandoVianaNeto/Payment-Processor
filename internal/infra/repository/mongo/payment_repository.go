package mongo_infra

import (
	"context"
	"errors"
	"math"
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
	filters := bson.D{
		{Key: "transactionProcessor", Value: bson.D{
			{Key: "$in", Value: bson.A{"default", "fallback"}},
		}},
	}

	requestedAt := bson.D{}

	if input.From != "" {
		requestedAt = append(requestedAt, bson.E{Key: "$gte", Value: input.From})
	}

	if input.To != "" {
		requestedAt = append(requestedAt, bson.E{Key: "$lte", Value: input.To})
	}

	if len(requestedAt) > 0 {
		filters = append(filters, bson.E{
			Key:   "requestedAt",
			Value: requestedAt,
		})
	}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filters}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$transactionProcessor"},
			{Key: "totalRequests", Value: bson.D{{Key: "$sum", Value: 1}}},
			{Key: "totalAmount", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "transactionProcessor", Value: "$_id"},
			{Key: "totalRequests", Value: 1},
			{Key: "totalAmount", Value: 1},
		}}},
	}

	cursor, err := f.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	response := &domain_response.PaymentSummaryResponse{}

	for _, result := range results {
		tp := result["transactionProcessor"].(string)
		totalRequests, _ := result["totalRequests"].(int32)
		totalAmount, _ := result["totalAmount"].(float64)

		switch tp {
		case "default":
			response.Default = domain_response.Summary{
				TotalRequests: int(totalRequests),
				TotalAmount:   roundToFixed(totalAmount, 1),
			}
		case "fallback":
			response.Fallback = domain_response.Summary{
				TotalRequests: int(totalRequests),
				TotalAmount:   roundToFixed(totalAmount, 1),
			}
		}
	}

	return response, nil
}

func (f *PaymentRepository) Delete(ctx context.Context, correlationId string) error {
	filter := bson.M{
		"correlationId": correlationId,
	}

	_, err := f.collection.DeleteOne(ctx, filter)

	return err
}

func roundToFixed(val float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Round(val*factor) / factor
}
