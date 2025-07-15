package server

import (
	"log"
	"net/http"
)

type Server struct {
	logger *log.Logger
	server http.Server
}

func NewServer(logger *log.Logger) *Server {

}
