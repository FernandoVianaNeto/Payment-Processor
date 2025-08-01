package web

import (
	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine, server *Server) *gin.Engine {
	{
		payments := engine.Group("/payments")
		{
			payments.POST("/", server.PaymentsHandler)
			payments.GET("/summary", server.GetSummaryHandler)
		}
	}

	return engine
}
