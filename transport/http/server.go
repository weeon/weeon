package http

import "github.com/gin-gonic/gin"

type Server struct {
	engine *gin.Engine
	router func(r *gin.Engine)
}

func NewGinServer(f func(r *gin.Engine)) *Server {
	return &Server{
		engine: gin.Default(),
		router: f,
	}
}

func (s *Server) Run() error {
	s.router(s.engine)
	return s.engine.Run(":8080")
}
