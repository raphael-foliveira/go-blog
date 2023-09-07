package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/blog-backend/pkg/res"
)

type healthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheck() (string, *chi.Mux) {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res.JSON(w, http.StatusOK, healthCheckResponse{"ok"})
	})
	return "/", router
}
