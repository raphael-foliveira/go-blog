package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/blog-backend/pkg/controller"
)

func Author(controller *controller.Author) *chi.Mux {
	router := chi.NewRouter()
	return router
}
