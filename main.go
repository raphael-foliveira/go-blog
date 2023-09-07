package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/raphael-foliveira/blog-backend/pkg/controller"
	"github.com/raphael-foliveira/blog-backend/pkg/database"
	"github.com/raphael-foliveira/blog-backend/pkg/repository"
	"github.com/raphael-foliveira/blog-backend/pkg/routes"
	"github.com/raphael-foliveira/blog-backend/pkg/server"
	"github.com/raphael-foliveira/blog-backend/pkg/service"
)

func main() {
	godotenv.Load()
	router := startRoutes()
	appServer := server.New(":3000", router)
	err := appServer.Run()
	if err != nil {
		panic(err)
	}
}

func attachMiddleware(r *chi.Mux) {
	log.Println("attaching middleware")
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
	}))
	r.Use(middleware.Logger)
}

func startRoutes() *chi.Mux {
	router := chi.NewRouter()
	attachMiddleware(router)
	injectDependencies(router)
	return router
}

func injectDependencies(router *chi.Mux) {
	database := database.MustConnect()
	postRepository := repository.NewPostRepository(database)
	authorRepository := repository.NewAuthorRepository(database)
	authorService := service.NewAuthorService(authorRepository, postRepository)
	postService := service.NewPostService(postRepository, authorRepository)
	authorController := controller.NewAuthorController(authorService)
	postController := controller.NewPostController(postService)
	router.Mount(routes.HealthCheck())
	router.Mount(routes.Author(authorController))
	router.Mount(routes.Post(postController))
}
