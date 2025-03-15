package handler

import (
	"github.com/matryer/way"
	"github.com/stymsinghss/Tweety/internal/service"
	"net/http"
)

type handler struct {
	*service.Service
}

func New(svc *service.Service) http.Handler {
	handler := handler{
		svc,
	}
	api := way.NewRouter()
	api.HandleFunc("GET", "/login", handler.loginUser)
	api.HandleFunc("PUT", "/register", handler.registerUser)

	r := way.NewRouter()
	r.Handle("*", "/api...", http.StripPrefix("/api", api))
	return r
}