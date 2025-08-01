package web

import (
	"context"

	gin "github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(
	ctx context.Context,
) *Server {
	router := gin.Default()

	server := &Server{}
	server.router = Routes(router, server)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
