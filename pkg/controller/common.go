package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/blog-backend/pkg/res"
)

func parseId(w http.ResponseWriter, r *http.Request) int64 {
	id := chi.URLParam(r, "id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		res.New(w).BadRequestError("invalid id")
		return 0
	}
	return intId
}
