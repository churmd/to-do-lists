package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/churmd/to-do-lists/internal/adapters/inbound"
)

func main() {

	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(jsonLogger)

	slog.Info("starting server")

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: inbound.FullHandler(),
	}

	err := server.ListenAndServe()
	if err != nil {
		slog.Error(fmt.Sprintf("server exited with error: %s", err.Error()))
	}
}
