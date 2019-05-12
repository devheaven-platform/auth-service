package controllers

import (
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// CreateMetricsController is used to create the metrics
// controller. It returns an instance of a chi.Mux router.
func CreateMetricsController() *chi.Mux {
	router := chi.NewRouter()
	router.Handle("/", promhttp.Handler())

	return router
}
