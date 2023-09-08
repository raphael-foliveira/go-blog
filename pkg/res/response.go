package res

import (
	"encoding/json"
	"net/http"
)

type response struct {
	w http.ResponseWriter
}

func New(w http.ResponseWriter) *response {
	return &response{w: w}
}

func (r *response) Status(status int) *response {
	r.w.WriteHeader(status)
	return r
}

func (r *response) JSON(body interface{}) error {
	return json.NewEncoder(r.w).Encode(body)
}

func (r *response) SendStatus(status int) error {
	r.w.WriteHeader(status)
	return nil
}

func (r *response) NotFoundError(message string) error {
	return r.Status(http.StatusNotFound).JSON(httpError{message})
}

func (r *response) BadRequestError(message string) error {
	return r.Status(http.StatusBadRequest).JSON(httpError{message})
}

func (r *response) InternalServerError(message string) error {
	return r.Status(http.StatusInternalServerError).JSON(httpError{message})
}

func (r *response) ConflictError(message string) error {
	return r.Status(http.StatusConflict).JSON(httpError{message})
}
