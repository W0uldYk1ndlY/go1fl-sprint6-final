package server

import (
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger *log.Logger
	Server *http.Server
}

func NewServer(logger *log.Logger) *Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/", handlers.HandleMain)
	handler.HandleFunc("/upload", handlers.HandleUpload)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ErrorLog:     logger,
		ReadTimeout:  5,
		WriteTimeout: 10,
		IdleTimeout:  15}

	newServer := Server{
		logger: logger,
		Server: server}

	return &newServer
}
