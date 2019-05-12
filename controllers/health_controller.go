package controllers

import (
	"net/http"

	"github.com/devheaven-platform/auth-service/models"
	"github.com/devheaven-platform/auth-service/utils"
	"github.com/go-chi/chi"
)

// CreateHealthController is used to create the health
// controller. It returns an instance of a chi.Mux router.
func CreateHealthController() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getHealth)

	return router
}

// getHealth handles the /health/ route. It takes an
// http.ResponseWriter and http.Request as parameters.
func getHealth(res http.ResponseWriter, req *http.Request) {
	utils.RespondJSON(res, http.StatusOK, models.Health{
		Message: "Service is running",
	})
}
