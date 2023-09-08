package routes

import (
	"net/http"

	"github.com/raphael-foliveira/blog-backend/pkg/res"
)

func wrapHandler(fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			res.New(w).InternalServerError(err.Error())
		}
	}
}
