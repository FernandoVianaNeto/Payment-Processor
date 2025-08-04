package cli

import (
	"context"

	configs "payment-gateway/cmd/config"
	payment_usecase "payment-gateway/internal/application/usecase/payments"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
	mongo_infra "payment-gateway/internal/infra/repository/mongo"
	"payment-gateway/internal/infra/web"
	mongoPkg "payment-gateway/pkg/mongo"
	natsclient "payment-gateway/pkg/nats"

	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	UseCases UseCases
}
type UseCases struct {
	CreatePaymentRequestUsecase domain_payment_usecase.CreatePaymentRequestUsecaseInterface
	GetPaymentsSummaryUsecase   domain_payment_usecase.GetPaymentsSummaryUsecaseInterface
}

type Repositories struct {
	PaymentsRepository domain_repository.PaymentRepositoryInterface
}

func NewApplication() *web.Server {
	ctx := context.Background()

	queueClient := natsclient.New(configs.NatsCfg.Host)
	queueClient.Connect()

	// redisClient := redis_client.InitRedis()

	mongoConnectionInput := mongoPkg.MongoInput{
		DSN:      configs.MongoCfg.Dsn,
		Database: configs.MongoCfg.Database,
	}

	db := mongoPkg.NewMongoDatabase(ctx, mongoConnectionInput)

	repositories := NewRepositories(ctx, db)

	usecases := NewUseCases(ctx, repositories.PaymentsRepository, queueClient)

	srv := web.NewServer(
		ctx,
		usecases.CreatePaymentRequestUsecase,
		usecases.GetPaymentsSummaryUsecase,
	)

	return srv
}

func NewRepositories(
	ctx context.Context,
	db *mongo.Database,
) Repositories {
	paymentsRepository := mongo_infra.NewPaymentRepository(db)

	return Repositories{
		PaymentsRepository: paymentsRepository,
	}
}

func NewUseCases(
	ctx context.Context,
	paymentRepository domain_repository.PaymentRepositoryInterface,
	queue *natsclient.NatsClient,
) UseCases {
	createPaymentRequestUsecase := payment_usecase.NewCreatePaymentRequestUsecase(paymentRepository, queue)
	getPaymentsSummaryUsecase := payment_usecase.NewGetPaymentsSummaryUsecase(paymentRepository)

	return UseCases{
		CreatePaymentRequestUsecase: createPaymentRequestUsecase,
		GetPaymentsSummaryUsecase:   getPaymentsSummaryUsecase,
	}
}
