package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

	appMiddleware "github.com/dadadam/go-http-example/api/middleware"
	authorRouter "github.com/dadadam/go-http-example/api/router/author"
	bookRouter "github.com/dadadam/go-http-example/api/router/book"
)

type Response struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	resp.Status = http.StatusOK
	return nil
}

func Init() *chi.Mux {
	r := chi.NewRouter()

	// Add middlewares
	// r.Use(middleware.Logger)
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	r.Use(appMiddleware.Logger("sjknvlksv", logger))
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

	r.Get("/health", healthCheck)

	return r
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	payload := make(map[string]interface{})
	payload["health"] = "ok"

	resp := &Response{Data: payload}
	render.Render(w, r, resp)
}
