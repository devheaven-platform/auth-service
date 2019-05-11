package controllers

import (
	"net/http"

	"github.com/devheaven-platform/auth-service/models"
	"github.com/devheaven-platform/auth-service/utils"
	"github.com/go-chi/chi"
)

// healthController represents the health controller and
// is used to group all the health endpoints together.
type healthController struct{}

// CreateHealthController is used to create the health
// controller. It returns an instance of a chi.Mux router.
func CreateHealthController() *chi.Mux {
	// Create the controller
	controller := &healthController{}

	// Create routes
	router := chi.NewRouter()
	router.Get("/", controller.getHealth)

	// Return the router
	return router
}

// getHealth handles the /health/ route. It takes an
// http.ResponseWriter and http.Request as parameters.
func (controller *healthController) getHealth(res http.ResponseWriter, req *http.Request) {
	utils.RespondJSON(res, http.StatusOK, models.Health{
		Message: "Service is running",
	})
}
