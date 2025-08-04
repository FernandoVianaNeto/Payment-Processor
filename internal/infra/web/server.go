package web

import (
	"context"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"

	gin "github.com/gin-gonic/gin"
)

type Server struct {
	router   *gin.Engine
	Usecases UseCases
}

type UseCases struct {
	CreatePaymentUsecase domain_payment_usecase.CreatePaymentRequestUsecaseInterface
}

func NewServer(
	ctx context.Context,
	createPaymentUsecase domain_payment_usecase.CreatePaymentRequestUsecaseInterface,
) *Server {
	router := gin.Default()

	server := &Server{
		router: router,
		Usecases: UseCases{
			CreatePaymentUsecase: createPaymentUsecase,
		},
	}
	server.router = Routes(router, server)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
