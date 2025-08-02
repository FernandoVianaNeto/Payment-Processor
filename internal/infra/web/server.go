package web

import (
	"context"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"

	gin "github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

type UseCases struct {
	CreatePaymentUsecase domain_payment_usecase.CreatePaymentUsecaseInterface
}

func NewServer(
	ctx context.Context,
	createPaymentUsecase domain_payment_usecase.CreatePaymentUsecaseInterface,
) *Server {
	router := gin.Default()

	server := &Server{}
	server.router = Routes(router, server)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
