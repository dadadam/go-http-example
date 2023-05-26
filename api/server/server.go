package server

import (
	"net/http"

	"github.com/dadadam/go-http-example/api/router"
)

type Server struct{}

func (s *Server) Start(addr string) {
	http.ListenAndServe(addr, router.Init())
}
