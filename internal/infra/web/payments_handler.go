package web

import (
	"net/http"
	"payment-gateway/internal/domain/dto"
	"payment-gateway/internal/infra/web/requests"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreatePaymentHandler(ctx *gin.Context) {
	var createRequest requests.CreatePaymentRequest

	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. Could not parse body"})
	}

	err := s.Usecases.CreatePaymentUsecase.Execute(ctx, dto.CreatePaymentDto{
		CorrelationId: createRequest.CorrelationId,
		Amount:        createRequest.Amount,
	})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Could not process payment request"})
	}

	ctx.Status(http.StatusOK)
}

func (s *Server) GetSummaryHandler(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
