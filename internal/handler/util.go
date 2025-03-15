package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respond(w http.ResponseWriter, data interface{}, statusCode int) {
	b, err := json.Marshal(data)
	if err != nil {
		respondInternalError(w, fmt.Errorf("could not marhal response: %v\n", err))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(b)
}

func respondInternalError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func respondError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, err.Error(), statusCode)
}
