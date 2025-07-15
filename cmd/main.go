package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "LOG:", 0)
	srv := server.NewServer(logger)

	if err := srv.Server.ListenAndServe(); err != nil {
		log.Fatalf("error while starting the server: %s", err.Error())
		return
	}

}
