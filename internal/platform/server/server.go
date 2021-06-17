package server

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"nequi.com/poc-services/internal/platform/server/handler/health"
	"nequi.com/poc-services/internal/platform/server/handler/credits"
	"nequi.com/poc-services/internal/services/credits"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
	getCreditsByClientService  services.GetCreditsByClientService
}

func New(host string, port uint, getCreditsByClientService services.GetCreditsByClientService ) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		getCreditsByClientService: getCreditsByClientService,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/creditsbyclient", credits.GetCreditsByClientHandler(s.getCreditsByClientService))
}
