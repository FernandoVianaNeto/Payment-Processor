package web

import (
	"context"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"

	gin "github.com/gin-gonic/gin"
)

type Server struct {
	router                    *gin.Engine
	CreatePaymentUsecase      domain_payment_usecase.CreatePaymentRequestUsecaseInterface
	GetPaymentsSummaryUsecase domain_payment_usecase.GetPaymentsSummaryUsecaseInterface
}

func NewServer(
	ctx context.Context,
	createPaymentUsecase domain_payment_usecase.CreatePaymentRequestUsecaseInterface,
	getPaymentsSummaryUsecase domain_payment_usecase.GetPaymentsSummaryUsecaseInterface,
) *Server {
	router := gin.Default()

	server := &Server{
		router:                    router,
		CreatePaymentUsecase:      createPaymentUsecase,
		GetPaymentsSummaryUsecase: getPaymentsSummaryUsecase,
	}
	server.router = Routes(router, server)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
