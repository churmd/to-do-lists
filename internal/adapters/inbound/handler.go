package inbound

import (
	"net/http"

	"github.com/churmd/to-do-lists/internal/adapters/inbound/rest"
	"github.com/go-chi/chi/v5"
)

func FullHandler() http.Handler {
	strictServer := rest.NewStrictHandler(rest.StrictChiServer{}, nil)
	router := chi.NewRouter()

	return rest.HandlerFromMux(strictServer, router)
}
