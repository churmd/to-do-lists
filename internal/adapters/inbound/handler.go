package inbound

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func FullHandler() http.Handler {
	strictServer := NewStrictHandler(StrictChiServer{}, nil)
	router := chi.NewRouter()

	return HandlerFromMux(strictServer, router)
}
