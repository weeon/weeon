package http

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Server struct {
	service string
	engine  *gin.Engine
	router  func(r *gin.Engine)
}

func NewGinServer(f func(r *gin.Engine), service string) *Server {
	return &Server{
		engine: gin.Default(),
		router: f,
	}
}

func (s *Server) Run() error {
	s.engine.Use(otelgin.Middleware(s.service))
	s.router(s.engine)
	return s.engine.Run(":8080")
}
