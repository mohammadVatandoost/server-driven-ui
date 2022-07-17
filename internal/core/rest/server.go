package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Engine *gin.Engine
	logger *logrus.Logger
	srv    *http.Server
	conf   Config
}

func (s *Server) Run() {
	port := s.conf.ListenPort
	addr := fmt.Sprintf(":%v", port)
	s.logger.Infof("CRM REST API Service Running, addr: %v \n", addr)

	s.srv = &http.Server{
		Addr:    addr,
		Handler: s.Engine,
	}

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Panicf("error on listen HTTP Server, error: %s", err.Error())
	}
}

func (s *Server) Shutdown(ctx context.Context) {
	if err := s.srv.Shutdown(ctx); err != nil {
		s.logger.Panicf("server forced to shutdown, err: %s", err.Error())
	}
}

func New(logger *logrus.Logger, conf Config) *Server {
	return &Server{
		logger: logger,
		Engine: gin.New(),
		conf:   conf,
	}
}
