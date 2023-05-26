package book

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func listBooks(w http.ResponseWriter, r *http.Request) {

}

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", listBooks)

	return r
}
