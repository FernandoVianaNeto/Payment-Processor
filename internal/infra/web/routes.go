package web

import (
	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine, server *Server) *gin.Engine {
	{
		payments := engine.Group("/")
		{
			payments.POST("/payments", server.CreatePaymentHandler)
			payments.GET("/payments-summary", server.GetSummaryHandler)
		}
	}

	return engine
}
