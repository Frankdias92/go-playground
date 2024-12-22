package main

import (
	"log/slog"
	"net/http"
	"time"

	api "github.com/frankdias92/go-playground/api-project"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		return
	}
	slog.Info("all system offline")
}

func run() error {
	handler := api.NewHandler()

	server := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8085",
		Handler:      handler,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
