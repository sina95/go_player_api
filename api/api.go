package api

import (
	"encoding/json"
	"net/http"
)

// balance parameters
type BalanceParams struct {
	Username string
}

// balance response
type BalanceResponse struct {
	// success code, usually 200
	Code int

	// player balance
	Balance int64
}

type Error struct {
	// error code
	Code int

	// message string
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error {
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured.", http.StatusInternalServerError)
	}
)