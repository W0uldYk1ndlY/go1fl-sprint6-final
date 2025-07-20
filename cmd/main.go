package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "LOG:", 0)
	srv := server.NewServer(logger)

	logger.Printf("Starting web-server at %s", srv.Server.Addr)
	os.Mkdir("result", 0755)

	if err := srv.Server.ListenAndServe(); err != nil {
		log.Fatalf("error while starting the server: %s", err.Error())
		return
	}

}
