package transport

import (
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// CreateTransport is used to create the metrics
// controller. It returns an instance of a chi.Mux router.
func CreateTransport() *chi.Mux {
	router := chi.NewRouter()
	router.Handle("/", promhttp.Handler())

	return router
}
