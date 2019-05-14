package transport

import (
	"net/http"

	"github.com/go-chi/chi"
)

// CreateTransport is used to create the swagger
// controller. It returns an instance of a chi.Mux router.
func CreateTransport() *chi.Mux {
	fs := http.FileServer(http.Dir("./dist"))

	router := chi.NewRouter()
	router.Handle("/*", http.StripPrefix("/docs/", fs))
	router.HandleFunc("/swagger.yaml", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./spec/swagger.yaml")
	})

	return router
}
