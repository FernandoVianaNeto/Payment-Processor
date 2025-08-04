package cli

import (
	"context"
	"net/http"

	configs "payment-gateway/cmd/config"
	payment_usecase "payment-gateway/internal/application/usecase/payments"
	domain_processors "payment-gateway/internal/domain/adapters/processors"
	"payment-gateway/internal/domain/adapters/queue"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
	"payment-gateway/internal/infra/adapter/processors"
	mongo_infra "payment-gateway/internal/infra/repository/mongo"
	http_client "payment-gateway/pkg/client/http"
	mongoPkg "payment-gateway/pkg/mongo"
	natsclient "payment-gateway/pkg/nats"

	"go.mongodb.org/mongo-driver/mongo"
)

type Worker struct {
	UseCases UseCases
}
type WorkerUseCases struct {
	ProcessPaymentRequestUsecase domain_payment_usecase.ProcessPaymentRequestUsecaseInterface
}

type WorkerRepositories struct {
	PaymentsRepository domain_repository.PaymentRepositoryInterface
}

type WorkerAdapters struct {
	ProcessorPaymentDefault  domain_processors.ProcessorsClientInterface
	ProcessorPaymentFallback domain_processors.ProcessorsClientInterface
}

type ConsumerInfra struct {
	ProcessorPaymentDefault  domain_processors.ProcessorsClientInterface
	ProcessorPaymentFallback domain_processors.ProcessorsClientInterface
	PaymentRepository        domain_repository.PaymentRepositoryInterface
	Queue                    *natsclient.NatsClient
	Usecases                 WorkerUseCases
}

func NewWorker() ConsumerInfra {
	ctx := context.Background()

	queueClient := natsclient.New(configs.NatsCfg.Host)
	queueClient.Connect()

	// redisClient := redis_client.InitRedis()

	mongoConnectionInput := mongoPkg.MongoInput{
		DSN:      configs.MongoCfg.Dsn,
		Database: configs.MongoCfg.Database,
	}

	db := mongoPkg.NewMongoDatabase(ctx, mongoConnectionInput)

	repositories := NewWorkerRepositories(ctx, db)

	adapters := NewWorkerAdapters(repositories.PaymentsRepository)

	usecases := NewWorkerUseCases(ctx, repositories.PaymentsRepository, adapters.ProcessorPaymentDefault, adapters.ProcessorPaymentFallback, queueClient)

	return ConsumerInfra{
		ProcessorPaymentDefault:  adapters.ProcessorPaymentDefault,
		ProcessorPaymentFallback: adapters.ProcessorPaymentFallback,
		PaymentRepository:        repositories.PaymentsRepository,
		Queue:                    queueClient,
		Usecases:                 usecases,
	}
}

func NewWorkerAdapters(
	paymentRepository domain_repository.PaymentRepositoryInterface,
) WorkerAdapters {
	paymentProcessorDefaultClient := http_client.NewBaseClient(
		configs.PaymentProcessorDefaultClientCfg.BaseUri,
		http.DefaultClient,
		configs.ApplicationCfg,
	)

	paymentProcessorFallbackClient := http_client.NewBaseClient(
		configs.PaymentProcessorFallbackClientCfg.BaseUri,
		http.DefaultClient,
		configs.ApplicationCfg,
		http_client.OptionalHeaders{
			Key: "Content-Type", Value: "application/json",
		},
	)

	paymentProcessorDefaultAdapter := processors.NewProcessorDefaultClient(paymentProcessorDefaultClient, paymentRepository)
	paymentProcessorFallbackAdapter := processors.NewProcessorDefaultClient(paymentProcessorFallbackClient, paymentRepository)

	return WorkerAdapters{
		ProcessorPaymentDefault:  paymentProcessorDefaultAdapter,
		ProcessorPaymentFallback: paymentProcessorFallbackAdapter,
	}
}

func NewWorkerRepositories(
	ctx context.Context,
	// redisClient *redis.Client,
	db *mongo.Database,
) Repositories {
	paymentsRepository := mongo_infra.NewPaymentRepository(db)

	return Repositories{
		PaymentsRepository: paymentsRepository,
	}
}

func NewWorkerUseCases(
	ctx context.Context,
	paymentRepository domain_repository.PaymentRepositoryInterface,
	processPaymentDefaultAdapter domain_processors.ProcessorsClientInterface,
	processPaymentFallbackAdapter domain_processors.ProcessorsClientInterface,
	queue queue.Client,
) WorkerUseCases {
	processPaymentRequestUsecase := payment_usecase.NewProcessPaymentRequestUsecase(paymentRepository, processPaymentDefaultAdapter, processPaymentFallbackAdapter, queue)

	return WorkerUseCases{
		ProcessPaymentRequestUsecase: processPaymentRequestUsecase,
	}
}
