package res

import "net/http"

type httpError struct {
	Error string `json:"error"`
}

func NotFound(w http.ResponseWriter, message string) error {
	return JSON(w, http.StatusNotFound, httpError{message})
}

func BadRequest(w http.ResponseWriter, message string) error {
	return JSON(w, http.StatusBadRequest, httpError{message})
}

func InternalServerError(w http.ResponseWriter, message string) error {
	return JSON(w, http.StatusInternalServerError, httpError{message})
}

func Conflict(w http.ResponseWriter, message string) error {
	return JSON(w, http.StatusConflict, httpError{message})
}
