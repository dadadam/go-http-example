package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	authorRouter "github.com/dadadam/go-http-example/api/router/author"
	bookRouter "github.com/dadadam/go-http-example/api/router/book"
)

func Init() *chi.Mux {
	r := chi.NewRouter()

	// Add middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Mount("/api/author", authorRouter.InitRouter())
	r.Mount("/api/book", bookRouter.InitRouter())

	return r
}
