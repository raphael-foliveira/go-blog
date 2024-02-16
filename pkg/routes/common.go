package routes

import (
	"errors"
	"net/http"

	"github.com/raphael-foliveira/blog-backend/pkg/controller"
	"github.com/raphael-foliveira/blog-backend/pkg/res"
)

func wrapHandler(fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			if errors.Is(err, controller.ErrParsingRequestBody) {
				res.Send(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
				return
			}
			res.Send(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}
