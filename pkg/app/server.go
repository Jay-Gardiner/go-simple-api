package app

import (
	"log"

	"github.com/Jay-Gardiner/go-simple-api/pkg/api"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	dataService api.DataService
}

func NewServer(router *gin.Engine, dataService api.DataService) *Server {
	return &Server{
		router:      router,
		dataService: dataService,
	}
}

func (s *Server) RunServer() error {
	// Initialise routes
	r := s.Routes()

	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error call RunServer on router: %v", err)
		return err
	}

	return nil
}
