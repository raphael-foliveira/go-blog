package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/blog-backend/pkg/controller"
)

func Post(controller *controller.Post) *chi.Mux {
	router := chi.NewRouter()
	return router
}
