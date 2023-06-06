package server

import (
	"fmt"
	"net/http"

	"github.com/dadadam/go-http-example/api/router"
	"github.com/dadadam/go-http-example/configs"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Config *configs.Config
}

func NewServer(cfg *configs.Config) *Server {
	server := &Server{Config: cfg}
	return server
}

func (s *Server) Start() {
	addr := fmt.Sprintf(":%d", s.Config.Port)

	log.WithFields(log.Fields{
		"module": "Server",
	}).Info(fmt.Sprintf("Http server starting on %s", addr))

	log.Fatal(http.ListenAndServe(addr, router.Init()))
}
