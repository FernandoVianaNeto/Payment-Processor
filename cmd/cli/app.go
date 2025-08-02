package cli

import (
	"context"

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
	CreatePaymentUsecase domain_payment_usecase.CreatePaymentUsecaseInterface
}

type Repositories struct {
	PaymentsRepository domain_repository.PaymentRepositoryInterface
}

func NewApplication() *web.Server {
	ctx := context.Background()

	queueClient := natsclient.New(configs.NatsCfg.Host)
	queueClient.Connect()

	redisClient := redis_client.InitRedis()

	repositories := NewRepositories(ctx, redisClient)

	usecases := NewUseCases(ctx, repositories.PaymentsRepository, queueClient)

	srv := web.NewServer(
		ctx,
		usecases.CreatePaymentUsecase,
	)

	return srv
}

func NewRepositories(
	ctx context.Context,
	redisClient *redis.Client,
) Repositories {
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
	createPaymentUsecase := payment_usecase.NewCreatePaymentUsecase(paymentRepository, queue)

	return UseCases{
		CreatePaymentUsecase: createPaymentUsecase,
	}
}
