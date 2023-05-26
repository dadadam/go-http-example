package author

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func listAuthors(w http.ResponseWriter, r *http.Request) {

}

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", listAuthors)

	return r
}
