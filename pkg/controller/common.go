package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func parseId(r *http.Request) (int64, error) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, ErrParsingId
	}
	return intId, nil
}

var ErrParsingId = errors.New("error parsing id from url")
