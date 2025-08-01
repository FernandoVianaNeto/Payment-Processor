package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) PaymentsHandler(ctx *gin.Context) {
	// if err != nil {
	// 	ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	// 	return
	// }

	ctx.Status(http.StatusOK)
}

func (s *Server) GetSummaryHandler(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
