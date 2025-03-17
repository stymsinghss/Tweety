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
	api.HandleFunc("POST", "/login", handler.loginUser)
	api.HandleFunc("POST", "/register", handler.registerUser)
	api.HandleFunc("GET", "/auth_user", handler.authUser)

	r := way.NewRouter()
	r.Handle("*", "/api...", http.StripPrefix("/api", handler.withAuth(api)))
	return r
}
