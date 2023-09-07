package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/blog-backend/pkg/controller"
)

func Author(controller *controller.Author) (string, *chi.Mux) {
	router := chi.NewRouter()
	router.Get("/", wrapHandler(controller.Find))
	router.Post("/", wrapHandler(controller.Create))
	router.Get("/{id}", wrapHandler(controller.FindOne))
	router.Put("/{id}", wrapHandler(controller.Update))
	router.Delete("/{id}", wrapHandler(controller.Delete))
	return "/authors", router
}
