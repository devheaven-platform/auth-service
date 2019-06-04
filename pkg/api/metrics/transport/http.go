package transport

import (
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// CreateHTTPTransport is used to create the metrics
// transport. It returns an instance of a chi.Mux
// router.
func CreateHTTPTransport() *chi.Mux {
	router := chi.NewRouter()
	router.Handle("/", promhttp.Handler())

	return router
}
