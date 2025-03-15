package handler

import (
	"encoding/json"
	"errors"
	"github.com/stymsinghss/Tweety/internal/service"
	"net/http"
)

type createUserInput struct {
	Email, Username string
}

// register -> handles register user functionality
func (h *handler) registerUser(w http.ResponseWriter, r *http.Request) {
	var in createUserInput
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.CreateUser(r.Context(), in.Email, in.Username)
	if errors.Is(err, service.ErrInvalidEmail) || errors.Is(err, service.ErrInvalidUsername) {
		respondError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if errors.Is(err, service.ErrEmailTaken) || errors.Is(err, service.ErrUsernameTaken) {
		respondError(w, err, http.StatusConflict)
		return
	}
	
	if err != nil {
		respondInternalError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}