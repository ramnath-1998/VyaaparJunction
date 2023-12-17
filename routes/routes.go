package routes

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ramnath.1998/vyaaparjunction/controllers"
)

func RunRoutes() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {

		r.Get("/hello", controllers.HelloWorld)
		r.Get("/products", controllers.GetProducts)
		r.Post("/products", controllers.PostProducts)
		r.Get("/products/categories", controllers.GetCategories)
		r.Post("/products/categories", controllers.PostProducts)

	})

	http.ListenAndServe(":4000", r)
	log.Printf("Routes started in port : 4000")
}
