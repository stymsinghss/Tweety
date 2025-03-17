package handler

import (
	"encoding/json"
	"errors"
	"github.com/stymsinghss/Tweety/internal/utils"
	"net/http"
)

type loginInput struct {
	Email string
}

// login -> handles login functionality
func (h *handler) loginUser(w http.ResponseWriter, r *http.Request) {
	var in loginInput
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call service layer
	out, err := h.Login(r.Context(), in.Email)

	// Check errors
	if errors.Is(err, utils.ErrUserNotFound) {
		respondError(w, err, http.StatusNotFound)
		return
	}

	if errors.Is(err, utils.ErrInvalidEmail) {
		respondError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		respondInternalError(w, err)
		return
	}

	respond(w, out, http.StatusOK)
}

// authUser -> get authenticated user from the token
func (h *handler) authUser(w http.ResponseWriter, r *http.Request) {
	// Call service layer
	user, err := h.AuthUser(r.Context())

	// Check errors
	if errors.Is(err, utils.ErrUnauthenticated) {
		respondError(w, err, http.StatusUnauthorized)
		return
	}
	if errors.Is(err, utils.ErrUserNotFound) {
		respondError(w, err, http.StatusNotFound)
		return
	}

	if err != nil {
		respondInternalError(w, err)
		return
	}

	respond(w, user, http.StatusOK)
}