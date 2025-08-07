package cli

import (
	"context"
	"fmt"

	configs "payment-gateway/cmd/config"
	payment_usecase "payment-gateway/internal/application/usecase/payments"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
	redis_payment_repository "payment-gateway/internal/infra/repository/redis/payments"
	"payment-gateway/internal/infra/web"
	natsclient "payment-gateway/pkg/nats"
	redis_client "payment-gateway/pkg/redis"

	"github.com/redis/go-redis/v9"
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
	err := queueClient.Connect()

	if err != nil {
		fmt.Println("❌ COULD NOT CONNECT WITH NATS:", err)
		panic(err)
	}

	fmt.Println("✅ SUCCESSFULLY CONNECTED WITH NATS")
	redisClient := redis_client.InitRedis()

	// mongoConnectionInput := mongoPkg.MongoInput{
	// 	DSN:      configs.MongoCfg.Dsn,
	// 	Database: configs.MongoCfg.Database,
	// }

	// db := mongoPkg.NewMongoDatabase(ctx, mongoConnectionInput)

	repositories := NewRepositories(ctx, redisClient)

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
	// db *mongo.Database,
	redisClient *redis.Client,
) Repositories {
	// paymentsRepository := mongo_infra.NewPaymentRepository(db)
	paymentsRepository := redis_payment_repository.NewPaymentsRepository(redisClient)

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
