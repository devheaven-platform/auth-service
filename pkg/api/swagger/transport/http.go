package transport

import (
	"net/http"

	"github.com/go-chi/chi"
)

// CreateHTTPTransport is used to create the swagger
// transport layer. It returns an instance of a chi.Mux
// router.
func CreateHTTPTransport() *chi.Mux {
	fs := http.FileServer(http.Dir("./dist"))

	router := chi.NewRouter()
	router.Handle("/*", http.StripPrefix("/docs/", fs))
	router.HandleFunc("/openapi.yaml", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./spec/openapi.yaml")
	})

	return router
}
