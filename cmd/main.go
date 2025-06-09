package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "morse-server ", log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(logger)

	logger.Println("Server is starting at :8080")
	err := srv.HTTP.ListenAndServe()
	if err != nil {
		logger.Fatal("Server failed:", err)
	}
}
